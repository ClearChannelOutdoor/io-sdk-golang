package photos

import "time"

type Photo struct {
	Angle       string         `json:"angle,omitempty" bson:"angle,omitempty"`
	BookingID   string         `json:"bookingID,omitempty" bson:"bookingID,omitempty"`
	CapturedAt  *time.Time     `json:"capturedAt" bson:"capturedAt"`
	CreatedAt   *time.Time     `json:"createdAt" bson:"createdAt"`
	CreativeID  string         `json:"creativeID,omitempty" bson:"creativeID,omitempty"`
	DisplayID   string         `json:"displayID" bson:"displayID"`
	ExternalIDs []string       `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string         `json:"photoID" bson:"photoID"`
	Status      *PhotoStatus   `json:"status,omitempty" bson:"status,omitempty"`
	Type        string         `json:"type" bson:"type"`
	UpdatedAt   *time.Time     `json:"updatedAt" bson:"updatedAt"`
	Urls        map[string]any `json:"urls" bson:"urls"`
}

type PhotoStatus struct {
	Featured    *bool `json:"featured,omitempty" bson:"featured,omitempty"`
	PopApproved *bool `json:"popApproved,omitempty" bson:"popApproved,omitempty"`
	Public      *bool `json:"public,omitempty" bson:"public,omitempty"`
}
