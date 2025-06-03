package networks

import "time"

type Address struct {
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Street1 string `json:"street1,omitempty" bson:"street1,omitempty"`
	Street2 string `json:"street2,omitempty" bson:"street2,omitempty"`
	Street3 string `json:"street3,omitempty" bson:"street3,omitempty"`
	Street4 string `json:"street4,omitempty" bson:"street4,omitempty"`
	Zip     string `json:"zip,omitempty" bson:"zip,omitempty"`
}
type GeoLocation struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
	} `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}
type HoursOfOperation struct {
	StartOffsetInMinutes int `json:"startOffset"`
	DurationInMinutes    int `json:"duration"`
}
type Location struct {
	Address     *Address     `json:"address,omitempty" bson:"address,omitempty"`
	GeoLocation *GeoLocation `json:"geolocation,omitempty" bson:"geolocation,omitempty"`
	Region      string       `json:"region,omitempty" bson:"region,omitempty"`
	Timezone    string       `json:"timezone,omitempty" bson:"timezone,omitempty"`
}

type Market struct {
	ExternalIDs               []string                         `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID                        string                           `json:"marketID,omitempty" bson:"marketID,omitempty"`
	CoreBasedStatisticalAreas []MarketCoreBasedStatisticalArea `json:"coreBasedStatisticalAreas,omitempty" bson:"coreBasedStatisticalAreas,omitempty"`
	DesignatedMarketAreas     []MarketDesignatedMarketArea     `json:"designatedMarketAreas,omitempty" bson:"designatedMarketAreas,omitempty"`
	Name                      string                           `json:"name,omitempty" bson:"name,omitempty"`
}

type MarketCoreBasedStatisticalArea struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type MarketDesignatedMarketArea struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type Network struct {
	ID              string                 `json:"networkID" bson:"networkID"`
	CreatedAt       *time.Time             `json:"createdAt,omitempty" bson:"createdAt"`
	Description     string                 `json:"description,omitempty" bson:"description,omitempty"`
	Digital         *NetworkDigital        `json:"digital,omitempty" bson:"digital,omitempty"`
	DisplayCount    *int                   `json:"displayCount,omitempty" bson:"-"`
	ExternalIDs     []string               `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	GeoLocation     *GeoLocation           `json:"geoLocation,omitempty" bson:"geoLocation,omitempty"`
	Market          *Market                `json:"market,omitempty" bson:"market,omitempty"`
	Name            string                 `json:"name,omitempty" bson:"name,omitempty"`
	Status          *NetworkStatus         `json:"status,omitempty" bson:"status,omitempty"`
	SynchedAt       time.Time              `json:"-" bson:"synchedAt,omitempty"`
	Type            string                 `json:"type,omitempty" bson:"type,omitempty"`
	UpdatedAt       *time.Time             `json:"updatedAt,omitempty" bson:"updatedAt"`
	Venue           *Venue                 `json:"venue,omitempty" bson:"venue,omitempty"`
	VisibleCreative *NetworkVisualCreative `json:"visibleCreative,omitempty" bson:"visibleCreative,omitempty"`
}
type NetworkDigital struct {
	EstimatedImpressionsPerSlot float32            `json:"estimatedImpressionsPerSlot,omitempty" bson:"estimatedImpressionsPerSlot,omitempty"`
	ExternalIDs                 []string           "json:\"externalIDs,omitempty\" bson:\"externalIDs,omitempty\""
	Hours                       []HoursOfOperation `json:"hours,omitempty" bson:"hours,omitempty"`
	Height                      int                `json:"height,omitempty" bson:"height,omitempty"`
	SlotCount                   int                `json:"slotCount,omitempty" bson:"slotCount,omitempty"`
	SlotDuration                int                `json:"slotDuration,omitempty" bson:"slotDuration,omitempty"`
	Width                       int                `json:"width,omitempty" bson:"width,omitempty"`
}
type NetworkDisplay struct {
	CreatedAt   *time.Time            `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	DisplayID   string                `json:"displayID,omitempty" bson:"displayID,omitempty"`
	ExternalIDs []string              `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	GeoLocation *GeoLocation          `json:"geoLocation,omitempty" bson:"geoLocation,omitempty"`
	NetworkID   string                `json:"networkID,omitempty" bson:"networkID,omitempty"`
	Status      *NetworkDisplayStatus `json:"status,omitempty" bson:"status,omitempty"`
	SynchedAt   time.Time             `json:"-" bson:"synchedAt,omitempty"`
	UpdatedAt   *time.Time            `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
type NetworkDisplayStatus struct {
	Active           *bool      `json:"active,omitempty" bson:"active,omitempty"`
	QuattroUpdatedAt *time.Time `json:"quattroUpdatedAt,omitempty" bson:"quattroUpdatedAt,omitempty"`
}

type NetworkStatus struct {
	Active           *bool      `json:"active,omitempty" bson:"active,omitempty"`
	QuattroUpdatedAt *time.Time `json:"quattroUpdatedAt,omitempty" bson:"quattroUpdatedAt,omitempty"`
}

type NetworkVisualCreative struct {
	Height         float64  `json:"height,omitempty" bson:"height,omitempty"`
	SupportedMedia []string `json:"supportedMedia,omitempty" bson:"supportedMedia,omitempty"`
	UnitOfMeasure  string   `json:"unitOfMeasure,omitempty" bson:"unitOfMeasure,omitempty"`
	Width          float64  `json:"width,omitempty" bson:"width,omitempty"`
}
type Venue struct {
	AirportCode string    `json:"airportCode,omitempty" bson:"airportCode,omitempty"`
	Division    string    `json:"division,omitempty" bson:"division,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string    `json:"venueID,omitempty" bson:"venueID,omitempty"`
	Location    *Location `json:"location,omitempty" bson:"location,omitempty"`
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Type        string    `json:"type,omitempty" bson:"type,omitempty"`
}
