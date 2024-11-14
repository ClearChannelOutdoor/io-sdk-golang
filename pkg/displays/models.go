package displays

import (
	"time"
)

type Display struct {
	AllowLocationBuy *bool                    `json:"allowLocationBuy,omitempty"`
	Artery           *DisplayArtery           `json:"artery,omitempty"`
	CreatedAt        *time.Time               `json:"createdAt,omitempty"`
	Description      string                   `json:"description,omitempty"`
	Digital          *MediaProductDigitalInfo `json:"digital,omitempty"`
	ExternalIDs      []string                 `json:"externalIDs,omitempty"`
	Facing           string                   `json:"facing,omitempty"`
	GeoLocation      *GeoLocation             `json:"geolocation,omitempty"`
	Height           *DisplayHeight           `json:"height,omitempty"`
	ID               string                   `json:"displayID"`
	Illumination     *DisplayIllumination     `json:"illumination,omitempty"`
	Market           *DisplayMarket           `json:"market,omitempty"`
	MediaProducts    []*MediaProduct          `json:"mediaProducts,omitempty"`
	ReadDirection    *string                  `json:"readDirection,omitempty"`
	StreetSide       string                   `json:"streetSide,omitempty"`
	Structure        *DisplayStructure        `json:"structure,omitempty"`
	Title            string                   `json:"title,omitempty"`
	UpdatedAt        *time.Time               `json:"updatedAt,omitempty"`
}

type DisplayArtery struct {
	Primary   string `json:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty"`
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
	CBSACode    string   `json:"cbsaCode,omitempty"`
	CBSAName    string   `json:"cbsaName,omitempty"`
	DMACode     string   `json:"dmaCode,omitempty"`
	DMAName     string   `json:"dmaName,omitempty"`
	Name        string   `json:"name,omitempty"`
	Number      string   `json:"number,omitempty"`
}

type DisplayStructure struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
}

type Size struct {
	Height        float64 `json:"height,omitempty"`
	UnitOfMeasure string  `json:"unitOfMeasure,omitempty"`
	Width         float64 `json:"width,omitempty"`
}

type GeoLocation struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates,omitempty"`
	Type string `json:"type,omitempty"`
}
