package hypixel

import (
	"time"
)

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

func (u *UUID) Format() {
	*u = (*u)[:8] + "-" + (*u)[8:12] + "-" + (*u)[12:16] + "-" + (*u)[16:20] + "-" + (*u)[20:32]
}

func (d Date) Time() time.Time {
	return time.UnixMilli(int64(d))
}
