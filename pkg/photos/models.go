package photos

import "time"

type Photo struct {
	Angle       string         `json:"angle,omitempty"`
	BookingID   string         `json:"bookingID,omitempty"`
	CapturedAt  *time.Time     `json:"capturedAt"`
	CreatedAt   *time.Time     `json:"createdAt"`
	CreativeID  string         `json:"creativeID,omitempty"`
	DisplayID   string         `json:"displayID"`
	ExternalIDs []string       `json:"externalIDs,omitempty"`
	ID          string         `json:"photoID"`
	Status      *PhotoStatus   `json:"status,omitempty"`
	SubmittedBy string         `json:"submittedBy,omitempty"`
	Type        string         `json:"type"`
	UpdatedAt   *time.Time     `json:"updatedAt"`
	Urls        map[string]any `json:"urls"`
}

type PhotoStatus struct {
	Featured    *bool `json:"featured,omitempty"`
	PopApproved *bool `json:"popApproved,omitempty"`
	Public      *bool `json:"public,omitempty"`
}
