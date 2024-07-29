package markets

import "time"

type Market struct {
	ID                        string                           `json:"marketID" bson:"marketID"`
	Code                      string                           `json:"code" bson:"code"`
	CoreBasedStatisticalAreas []MarketCoreBasedStatisticalArea `json:"coreBasedStatisticalAreas,omitempty" bson:"coreBasedStatisticalAreas,omitempty"`
	CreatedAt                 time.Time                        `json:"createdAt" bson:"createdAt"`
	DesignatedMarketAreas     []MarketDesignatedMarketArea     `json:"designatedMarketAreas,omitempty" bson:"designatedMarketAreas,omitempty"`
	ExternalIDs               []string                         `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Location                  struct {
		City    string `json:"city,omitempty" bson:"city,omitempty"`
		Country string `json:"country,omitempty" bson:"country,omitempty"`
		State   string `json:"state,omitempty" bson:"state,omitempty"`
	} `json:"location,omitempty" bson:"location,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Timezone  string    `json:"timezone,omitempty" bson:"timezone,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type MarketCoreBasedStatisticalArea struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type MarketDesignatedMarketArea struct {
	Code string `json:"code,omitempty" bson:"code,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
