package markets

import "time"

type Location struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

type Market struct {
	ID          string    `json:"marketID"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	ExternalIDs []string  `json:"externalIDs,omitempty"`
	Location    Location  `json:"location,omitempty"`
	Name        string    `json:"name"`
	Timezone    string    `json:"timezone,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
