package hypixel

import "strings"

type (
	Rank             string
	UserRank         Rank
	DonatorRank      Rank
	SubscriptionRank Rank
	Player           struct {
		Cosmetics
		Times
		ID                string   `json:"_id"`
		UUID              UUID     `json:"uuid"`
		DisplayName       string   `json:"displayname,omitempty"`
		PlayerName        string   `json:"playername,omitempty"`
		Aliases           []string `json:"knownAliases"`
		Achievements      []string `json:"achievementsOneTime,omitempty"`
		AchievementPoints uint     `json:"achievementPoints"`
		Karma             uint     `json:"karma"`
		EXP               float64  `json:"networkExp"`
		Socials           struct {
			Links struct {
				Discord string `json:"DISCORD"`
				Twitter string `json:"TWITTER"`
				Forum   string `json:"HYPIXEL"`
			} `json:"links,omitempty"`
		} `json:"socialMedia,omitempty"`
		Rank             DonatorRank      `json:"newPackageRank,omitempty"`
		SubscriptionRank SubscriptionRank `json:"monthlyPackageRank,omitempty"`
		Stats            struct{}         `json:"-"`
		FriendRequests   []UUID           `json:"friendRequestsUuid,omitempty"`
	}
	Cosmetics struct {
		VanityFavourites VanityFavourites `json:"vanityFavorites"`
		Gadget           string           `json:"currentGadget,omitempty"`
		ClickEffect      string           `json:"currentClickEffect,omitempty"`
		ParticlePack     string           `json:"ParticlePack,omitempty"`
		PlusColor        string           `json:"rankPlusColor,omitempty"`
		Cloak            string           `json:"currentCloak,omitempty"`
	}
	Times struct {
		FirstLogin Date `json:"firstLogin,omitempty"`
		LastLogin  Date `json:"lastLogin,omitempty"`
		LastLogout Date `json:"lastLogout,omitempty"`
	}
	VanityFavourites string
	PlayerStatus     struct {
		Online bool   `json:"online"`
		Game   string `json:"gameType"`
		Mode   string `json:"mode"`
		Map    string `json:"map"`
	}
)

func (v VanityFavourites) Parse() []string {
	return strings.Split(string(v), ";")
}
