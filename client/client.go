package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Gophigure/gopixel/httputil"
	"github.com/Gophigure/gopixel/hypixel"
	"net/http"
	"strconv"
	"time"
)

type (
	Client struct {
		Key        hypixel.APIKey
		Context    context.Context
		HttpClient httputil.Client
		Cabinet    *Cabinet
	}
	store struct {
		set   time.Time
		value interface{}
	}
	Cabinet struct {
		keyInfoStore map[hypixel.UUID]*store
		uuidStore    map[string]*store
		playerStore  map[hypixel.UUID]*store
		statusStore  map[hypixel.UUID]*store
	}
)

func New(ctx context.Context, key hypixel.APIKey, client httputil.Client) *Client {
	return &Client{
		Key:        key,
		Context:    ctx,
		HttpClient: client,
		Cabinet:    newCabinet(),
	}
}

func newCabinet() *Cabinet {
	return &Cabinet{
		keyInfoStore: make(map[hypixel.UUID]*store),
		uuidStore:    make(map[string]*store),
		playerStore:  make(map[hypixel.UUID]*store),
		statusStore:  make(map[hypixel.UUID]*store),
	}
}

func (c *Client) RequestJSON(v interface{}, method, url string) error {
	if v == nil {
		return nil
	}

	req, err := c.HttpClient.NewRequest(c.Context, method, url)
	if err != nil {
		return err
	}

	req.AddHeader(http.Header{
		"API-Key":      []string{string(c.Key)},
		"Content-Type": []string{"application/json"},
	})

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	body, status := res.GetBody(), res.GetStatus()
	defer body.Close()

	decoder := json.NewDecoder(body)

	if status < 200 || status > 299 {
		raw := struct {
			Cause string `json:"cause,omitempty"`
		}{}

		if err = decoder.Decode(&raw); err != nil {
			return err
		}

		if raw.Cause == "" {
			raw.Cause = "Unknown"
		}

		return errors.New("Received error response. Code: " + strconv.FormatInt(int64(status), 10) + " Cause: " + raw.Cause)
	}

	if err = decoder.Decode(&v); err != nil {
		return err
	}

	return nil
}

func (c *Client) NameToUUID(name string) (hypixel.UUID, error) {
	if uuid, ok := c.Cabinet.uuidStore[name]; ok && time.Since(uuid.set) < 24 * time.Hour {
		return uuid.value.(hypixel.UUID), nil
	}

	raw := struct {
		ID hypixel.UUID `json:"id"`
	}{}

	if err := c.RequestJSON(&raw, "GET", "https://api.mojang.com/users/profiles/minecraft/"+name); err != nil {
		return "", err
	}

	raw.ID.Format()
	c.Cabinet.uuidStore[name] = &store{
		set:   time.Now(),
		value: raw.ID,
	}

	return raw.ID, nil
}

func (c *Client) KeyInfo() (*hypixel.APIKeyInformation, error) {
	raw := struct {
		Success bool                      `json:"success"`
		KeyInfo hypixel.APIKeyInformation `json:"record"`
	}{}

	if err := c.RequestJSON(&raw, "GET", hypixel.BaseURL+"key"); err != nil {
		return nil, err
	}

	return &raw.KeyInfo, nil
}

func (c *Client) Player(uuid hypixel.UUID) (*hypixel.Player, error) {
	if player, ok := c.Cabinet.playerStore[uuid]; ok && time.Since(player.set) < time.Hour {
		return player.value.(*hypixel.Player), nil
	}

	raw := struct {
		Success bool           `json:"success"`
		Player  hypixel.Player `json:"player,omitempty"`
	}{}

	if err := c.RequestJSON(&raw, "GET", hypixel.BaseURL+"player?uuid="+string(uuid)); err != nil {
		return nil, err
	}

	c.Cabinet.playerStore[uuid] = &store{
		set:   time.Now(),
		value: &raw.Player,
	}

	return &raw.Player, nil
}

func (c *Client) PlayerStatus(uuid hypixel.UUID) (*hypixel.PlayerStatus, error) {
	if status, ok := c.Cabinet.statusStore[uuid]; ok && time.Since(status.set) > time.Hour {
		return status.value.(*hypixel.PlayerStatus), nil
	}

	raw := struct {
		Success bool                 `json:"success"`
		Status  hypixel.PlayerStatus `json:"session,omitempty"`
	}{}

	if err := c.RequestJSON(&raw, "GET", hypixel.BaseURL+"status?uuid="+string(uuid)); err != nil {
		return nil, err
	}

	c.Cabinet.statusStore[uuid] = &store{
		set: time.Now(),
		value: &raw.Status,
	}

	return &raw.Status, nil
}
