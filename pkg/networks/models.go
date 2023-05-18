package networks

import "time"

type Digital struct {
	EstimatedImpressionsPerSlot float32            `json:"estimatedImpressionsPerSlot,omitempty"`
	ExternalIDs                 []string           `json:"externalIDs,omitempty"`
	Hours                       []HoursOfOperation `json:"hours,omitempty"`
	Height                      int                `json:"height,omitempty"`
	SlotCount                   int                `json:"slotCount,omitempty"`
	SlotDuration                int                `json:"slotDuration,omitempty"`
	Width                       int                `json:"width,omitempty"`
}

type HoursOfOperation struct {
	StartOffsetInMinutes int `json:"startOffset"`
	DurationInMinutes    int `json:"duration"`
}

type Market struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"marketID,omitempty"`
	Name        string   `json:"name,omitempty"`
}

type Network struct {
	ID              string           `json:"networkID"`
	CreatedAt       *time.Time       `json:"createdAt,omitempty"`
	Description     string           `json:"description,omitempty"`
	Digital         *Digital         `json:"digital,omitempty"`
	DisplayCount    int              `json:"displayCount,omitempty"`
	ExternalIDs     []string         `json:"externalIDs,omitempty"`
	Market          *Market          `json:"market,omitempty"`
	Name            string           `json:"name,omitempty"`
	Type            string           `json:"type,omitempty"`
	UpdatedAt       *time.Time       `json:"updatedAt,omitempty"`
	Venue           *Venue           `json:"venue,omitempty"`
	VisibleCreative *VisibleCreative `json:"visibleCreative,omitempty"`
}

type NetworkDisplay struct {
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	DisplayID   string    `json:"displayID,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty"`
	NetworkID   string    `json:"networkID,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type Venue struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"venueID,omitempty"`
	Name        string   `json:"name,omitempty"`
}

type VisibleCreative struct {
	Height         float64  `json:"height,omitempty"`
	SupportedMedia []string `json:"supportedMedia,omitempty"`
	UnitOfMeasure  string   `json:"unitOfMeasure,omitempty"`
	Width          float64  `json:"width,omitempty"`
}
