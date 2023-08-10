package markets

import "time"

type Location struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

type Market struct {
	ID                        string                           `json:"marketID"`
	CoreBasedStatisticalAreas []MarketCoreBasedStatisticalArea `json:"coreBasedStatisticalAreas,omitempty" bson:"coreBasedStatisticalAreas,omitempty"`
	Code                      string                           `json:"code"`
	CreatedAt                 time.Time                        `json:"createdAt"`
	DesignatedMarketAreas     []MarketDesignatedMarketArea     `json:"designatedMarketAreas,omitempty" bson:"designatedMarketAreas,omitempty"`
	ExternalIDs               []string                         `json:"externalIDs,omitempty"`
	Location                  Location                         `json:"location,omitempty"`
	Name                      string                           `json:"name"`
	Timezone                  string                           `json:"timezone,omitempty"`
	UpdatedAt                 time.Time                        `json:"updatedAt"`
}

type MarketCoreBasedStatisticalArea struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type MarketDesignatedMarketArea struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
