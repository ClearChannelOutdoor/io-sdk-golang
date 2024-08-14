package displays

import (
	"time"
)

type Display struct {
	AllowLocationBuy *bool                    `json:"allowLocationBuy,omitempty" bson:"allowLocationBuy,omitempty"`
	Artery           *DisplayArtery           `json:"artery,omitempty" bson:"artery,omitempty"`
	CreatedAt        *time.Time               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Description      string                   `json:"description,omitempty" bson:"description,omitempty"`
	Digital          *MediaProductDigitalInfo `json:"digital,omitempty" bson:"digital,omitempty"` // TODO: after import is complete, set bson hint to "-"
	ExternalIDs      []string                 `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Facing           string                   `json:"facing,omitempty" bson:"facing,omitempty"`
	GeoLocation      *GeoLocation             `json:"geolocation,omitempty" bson:"geolocation,omitempty"`
	Height           *DisplayHeight           `json:"height,omitempty" bson:"height,omitempty"`
	ID               string                   `json:"displayID" bson:"displayID"`
	Illumination     *DisplayIllumination     `json:"illumination,omitempty" bson:"illumination,omitempty"`
	Market           *DisplayMarket           `json:"market,omitempty" bson:"market,omitempty"`
	MediaProducts    []*MediaProduct          `json:"mediaProducts,omitempty" bson:"mediaProducts,omitempty"`
	ReadDirection    *string                  `json:"readDirection,omitempty" bson:"readDirection,omitempty"`
	StreetSide       string                   `json:"streetSide,omitempty" bson:"streetSide,omitempty"`
	Structure        *DisplayStructure        `json:"structure,omitempty" bson:"structure,omitempty"`
	Title            string                   `json:"title,omitempty" bson:"title,omitempty"`
	UpdatedAt        *time.Time               `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type DisplayArtery struct {
	Primary   string `json:"primary,omitempty" bson:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty" bson:"secondary,omitempty"`
}

type DisplayHeight struct {
	FromGround    float64 `json:"fromGround,omitempty" bson:"fromGround,omitempty"`
	UnitOfMeasure string  `json:"unitOfMeasure,omitempty" bson:"unitOfMeasure,omitempty"`
}

type DisplayIllumination struct {
	Hours       int  `json:"hours,omitempty" bson:"hours,omitempty"`
	Illuminated bool `json:"illuminated,omitempty" bson:"illuminated,omitempty"`
}

type DisplayMarket struct {
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string   `json:"marketID,omitempty" bson:"marketID,omitempty"`
	CBSACode    string   `json:"cbsaCode,omitempty" bson:"cbsaCode,omitempty"`
	CBSAName    string   `json:"cbsaName,omitempty" bson:"cbsaName,omitempty"`
	DMACode     string   `json:"dmaCode,omitempty" bson:"dmaCode,omitempty"`
	DMAName     string   `json:"dmaName,omitempty" bson:"dmaName,omitempty"`
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
	Number      string   `json:"number,omitempty" bson:"number,omitempty"`
}

type DisplayStructure struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	// ID          string   `json:"structureID,omitempty"`
}

type Size struct {
	Height        float64 `json:"height,omitempty" bson:"height,omitempty"`
	UnitOfMeasure string  `json:"unitOfMeasure,omitempty" bson:"unitOfMeasure,omitempty"`
	Width         float64 `json:"width,omitempty" bson:"width,omitempty"`
}

type GeoLocation struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
	} `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}
