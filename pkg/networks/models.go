package networks

import "time"

type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	Street1 string `json:"street1,omitempty"`
	Street2 string `json:"street2,omitempty"`
	Street3 string `json:"street3,omitempty"`
	Street4 string `json:"street4,omitempty"`
	Zip     string `json:"zip,omitempty"`
}
type GeoLocation struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates,omitempty"`
	Type string `json:"type,omitempty"`
}
type HoursOfOperation struct {
	StartOffsetInMinutes int `json:"startOffset"`
	DurationInMinutes    int `json:"duration"`
}
type Location struct {
	Address     *Address     `json:"address,omitempty"`
	GeoLocation *GeoLocation `json:"geolocation,omitempty"`
	Region      string       `json:"region,omitempty"`
	Timezone    string       `json:"timezone,omitempty"`
}

type Market struct {
	ExternalIDs               []string                         `json:"externalIDs,omitempty"`
	ID                        string                           `json:"marketID,omitempty"`
	CoreBasedStatisticalAreas []MarketCoreBasedStatisticalArea `json:"coreBasedStatisticalAreas,omitempty"`
	DesignatedMarketAreas     []MarketDesignatedMarketArea     `json:"designatedMarketAreas,omitempty"`
	Name                      string                           `json:"name,omitempty"`
}

type MarketCoreBasedStatisticalArea struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type MarketDesignatedMarketArea struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Network struct {
	ID              string                 `json:"networkID"`
	CreatedAt       *time.Time             `json:"createdAt,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Digital         *NetworkDigital        `json:"digital,omitempty"`
	DisplayCount    *int                   `json:"displayCount,omitempty"`
	ExternalIDs     []string               `json:"externalIDs,omitempty"`
	GeoLocation     *GeoLocation           `json:"geoLocation,omitempty"`
	Market          *Market                `json:"market,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Status          *NetworkStatus         `json:"status,omitempty"`
	SynchedAt       time.Time              `json:"-"`
	Type            string                 `json:"type,omitempty"`
	UpdatedAt       *time.Time             `json:"updatedAt,omitempty"`
	Venue           *Venue                 `json:"venue,omitempty"`
	VisibleCreative *NetworkVisualCreative `json:"visibleCreative,omitempty"`
}
type NetworkDigital struct {
	EstimatedImpressionsPerSlot float32            `json:"estimatedImpressionsPerSlot,omitempty"`
	ExternalIDs                 []string           "json:\"externalIDs,omitempty\" bson:\"externalIDs,omitempty\""
	Hours                       []HoursOfOperation `json:"hours,omitempty"`
	Height                      int                `json:"height,omitempty"`
	SlotCount                   int                `json:"slotCount,omitempty"`
	SlotDuration                int                `json:"slotDuration,omitempty"`
	Width                       int                `json:"width,omitempty"`
}
type NetworkDisplay struct {
	CreatedAt   *time.Time            `json:"createdAt,omitempty"`
	DisplayID   string                `json:"displayID,omitempty"`
	ExternalIDs []string              `json:"externalIDs,omitempty"`
	GeoLocation *GeoLocation          `json:"geoLocation,omitempty"`
	NetworkID   string                `json:"networkID,omitempty"`
	Status      *NetworkDisplayStatus `json:"status,omitempty"`
	SynchedAt   time.Time             `json:"-"`
	UpdatedAt   *time.Time            `json:"updatedAt,omitempty"`
}
type NetworkDisplayStatus struct {
	Active           *bool      `json:"active,omitempty"`
	QuattroUpdatedAt *time.Time `json:"quattroUpdatedAt,omitempty"`
}

type NetworkStatus struct {
	Active           *bool      `json:"active,omitempty"`
	QuattroUpdatedAt *time.Time `json:"quattroUpdatedAt,omitempty"`
}

type NetworkVisualCreative struct {
	Height         float64  `json:"height,omitempty"`
	SupportedMedia []string `json:"supportedMedia,omitempty"`
	UnitOfMeasure  string   `json:"unitOfMeasure,omitempty"`
	Width          float64  `json:"width,omitempty"`
}
type Venue struct {
	AirportCode string    `json:"airportCode,omitempty"`
	Division    string    `json:"division,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty"`
	ID          string    `json:"venueID,omitempty"`
	Location    *Location `json:"location,omitempty"`
	Name        string    `json:"name,omitempty"`
	Type        string    `json:"type,omitempty"`
}
