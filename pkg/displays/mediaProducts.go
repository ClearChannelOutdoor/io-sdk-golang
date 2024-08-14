package displays

import (
	"time"
)

type DigitalHoursOfOperation struct {
	StartOffset       int `json:"startOffset" bson:"startOffset"`
	DurationInMinutes int `json:"duration" bson:"duration"`
}

type MediaProduct struct {
	Detail          string                       `json:"detail,omitempty" bson:"detail,omitempty"`
	Digital         *bool                        `json:"digital,omitempty" bson:"digital,omitempty"`
	DigitalInfo     *MediaProductDigitalInfo     `json:"digitalInfo,omitempty" bson:"digitalInfo,omitempty"`
	ExternalIDs     []string                     `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ProductID       string                       `json:"productID,omitempty" bson:"productID,omitempty"`
	Size            *Size                        `json:"size,omitempty" bson:"size,omitempty"`
	Status          *MediaProductStatus          `json:"status,omitempty" bson:"status,omitempty"`
	Type            string                       `json:"type,omitempty" bson:"type,omitempty"`
	VisibleCreative *MediaProductVisibleCreative `json:"visibleCreative,omitempty" bson:"visibleCreative,omitempty"`
}

type MediaProductDigitalInfo struct {
	Hours        []*DigitalHoursOfOperation `json:"hours,omitempty" bson:"hours,omitempty"`
	Height       int                        `json:"height,omitempty" bson:"height,omitempty"`
	SlotCount    int                        `json:"slotCount,omitempty" bson:"slotCount,omitempty"`
	SlotDuration int                        `json:"slotDuration,omitempty" bson:"slotDuration,omitempty"`
	Width        int                        `json:"width,omitempty" bson:"width,omitempty"`
}

type MediaProductStatus struct {
	Active     *bool                `json:"active,omitempty" bson:"active,omitempty"`
	FinishDate *time.Time           `json:"finishDate,omitempty" bson:"finishDate,omitempty"`
	LiveDate   *time.Time           `json:"liveDate,omitempty" bson:"liveDate,omitempty"`
	Saleable   *bool                `json:"saleable,omitempty" bson:"saleable,omitempty"`
	Sync       map[string]time.Time `json:"sync,omitempty" bson:"sync,omitempty"`
}

type MediaProductVisibleCreative struct {
	Height         float64          `json:"height,omitempty" bson:"height,omitempty"`
	SupportedMedia []SupportedMedia `json:"supportedMedia,omitempty" bson:"supportedMedia,omitempty"`
	UnitOfMeasure  string           `json:"unitOfMeasure,omitempty" bson:"unitOfMeasure,omitempty"`
	Width          float64          `json:"width,omitempty" bson:"width,omitempty"`
}

type SupportedMedia string
