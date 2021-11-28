package hypixel

import "time"

const BaseURL = "https://api.hypixel.net/"

type (
	UUID              string
	APIKey            UUID
	APIKeyInformation struct {
		Key          APIKey `json:"key"`
		Owner        UUID   `json:"owner,omitempty"`
		Limit        int    `json:"limit"`
		PastQueries  int    `json:"queriesInPastMin"`
		TotalQueries int    `json:"totalQueries"`
	}
	Date int64
)

func (d Date) Time() time.Time {
	return time.UnixMilli(int64(d))
}
