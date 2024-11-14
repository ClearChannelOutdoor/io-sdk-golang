package displays

import (
	"time"
)

type DigitalHoursOfOperation struct {
	StartOffset       int `json:"startOffset"`
	DurationInMinutes int `json:"duration"`
}

type MediaProduct struct {
	Code            *MediaProductCode            `json:"code,omitempty"`
	Detail          string                       `json:"detail,omitempty"`
	Digital         *bool                        `json:"digital,omitempty"`
	DigitalInfo     *MediaProductDigitalInfo     `json:"digitalInfo,omitempty"`
	ExternalIDs     []string                     `json:"externalIDs,omitempty"`
	ProductID       string                       `json:"productID,omitempty"`
	Size            *Size                        `json:"size,omitempty"`
	Status          *MediaProductStatus          `json:"status,omitempty"`
	Type            string                       `json:"type,omitempty"`
	VisibleCreative *MediaProductVisibleCreative `json:"visibleCreative,omitempty"`
}

type MediaProductCode struct {
	ProductCode string `json:"productCode,omitempty"`
	TypeCode    string `json:"typeCode,omitempty"`
}

type MediaProductDigitalInfo struct {
	Hours        []*DigitalHoursOfOperation `json:"hours,omitempty"`
	Height       int                        `json:"height,omitempty"`
	SlotCount    int                        `json:"slotCount,omitempty"`
	SlotDuration int                        `json:"slotDuration,omitempty"`
	Width        int                        `json:"width,omitempty"`
}

type MediaProductStatus struct {
	Active     *bool                `json:"active,omitempty"`
	FinishDate *time.Time           `json:"finishDate,omitempty"`
	LiveDate   *time.Time           `json:"liveDate,omitempty"`
	Saleable   *bool                `json:"saleable,omitempty"`
	Sync       map[string]time.Time `json:"sync,omitempty"`
}

type MediaProductVisibleCreative struct {
	Height         float64  `json:"height,omitempty"`
	SupportedMedia []string `json:"supportedMedia,omitempty"`
	UnitOfMeasure  string   `json:"unitOfMeasure,omitempty"`
	Width          float64  `json:"width,omitempty"`
}

type SupportedMedia string
