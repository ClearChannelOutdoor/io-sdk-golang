package displays

import (
	"time"
)

type Display struct {
	AllowLocationBuy *bool                `json:"allowLocationBuy,omitempty"`
	Artery           *DisplayArtery       `json:"artery,omitempty"`
	CreatedAt        *time.Time           `json:"createdAt,omitempty"`
	Description      string               `json:"description,omitempty"`
	Digital          *DisplayDigital      `json:"digital,omitempty"`
	ExternalIDs      []string             `json:"externalIDs,omitempty"`
	Facing           string               `json:"facing,omitempty"`
	FinishDate       *time.Time           `json:"finishDate,omitempty"`
	GeoLocation      *GeoLocation         `json:"geolocation,omitempty"`
	Height           *DisplayHeight       `json:"height,omitempty"`
	ID               string               `json:"displayID"`
	Illumination     *DisplayIllumination `json:"illumination,omitempty"`
	LiveDate         *time.Time           `json:"liveDate,omitempty"`
	Market           *DisplayMarket       `json:"market,omitempty"`
	MediaProducts    []MediaProduct       `json:"mediaProducts,omitempty"`
	ReadDirection    *string              `json:"readDirection,omitempty"`
	Route            *DisplayRoute        `json:"route,omitempty"`
	Saleable         *bool                `json:"saleable,omitempty"`
	StreetSide       string               `json:"streetSide,omitempty"`
	Structure        *DisplayStructure    `json:"structure,omitempty"`
	Title            string               `json:"title,omitempty"`
	UpdatedAt        *time.Time           `json:"updatedAt,omitempty"`
}

type DisplayArtery struct {
	Primary   string `json:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty"`
}

type DisplayDigital struct {
	ExternalIDs  []string           `json:"externalIDs,omitempty"`
	Hours        []HoursOfOperation `json:"hours,omitempty"`
	Height       int                `json:"height,omitempty"`
	SlotCount    int                `json:"slotCount,omitempty"`
	SlotDuration int                `json:"slotDuration,omitempty"`
	Width        int                `json:"width,omitempty"`
}

type DisplayHeight struct {
	FromGround    float64 `json:"fromGround,omitempty"`
	UnitOfMeasure string  `json:"unitOfMeasure,omitempty"`
}

type DisplayIllumination struct {
	Hours       int  `json:"hours,omitempty"`
	Illuminated bool `json:"illuminated,omitempty"`
}

type DisplayMarket struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"marketID,omitempty"`
	Name        string   `json:"name,omitempty"`
	Number      string   `json:"number,omitempty"`
}

type DisplayRoute struct {
	Name     string  `json:"name,omitempty"`
	Sequence float32 `json:"sequence,omitempty"`
}

type DisplayStructure struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"structureID,omitempty"`
}

type HoursOfOperation struct {
	StartOffset       int `json:"startOffset,omitempty"`
	DurationInMinutes int `json:"duration,omitempty"`
}

type MediaProduct struct {
	Detail      string   `json:"detail,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"productID,omitempty"`
	Size        struct {
		Height        float64 `json:"height,omitempty"`
		UnitOfMeasure string  `json:"unitOfMeasure,omitempty"`
		Width         float64 `json:"width,omitempty"`
	} `json:"size,omitempty"`
	Type            string `json:"type,omitempty"`
	VisibleCreative struct {
		Height         float64  `json:"height,omitempty"`
		SupportedMedia []string `json:"supportedMedia,omitempty"`
		UnitOfMeasure  string   `json:"unitOfMeasure,omitempty"`
		Width          float64  `json:"width,omitempty"`
	} `json:"visibleCreative,omitempty"`
}

type GeoLocation struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude,omitempty"`
		Longitude float64 `json:"longitude,omitempty"`
	} `json:"coordinates,omitempty"`
	Type string `json:"type,omitempty"`
}
