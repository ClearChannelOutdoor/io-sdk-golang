package pricing

import (
	"time"
)

type Quote struct {
	CreatedAt   *time.Time `json:"createdAt" bson:"createdAt"`
	ExternalIDs []string   `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	// includes
	// salesforce:opportunity:a3HVT000001HYBB2A4
	// salesforce:plan:a3JVT000001Kh532AC
	ID             string          `json:"quoteID" bson:"quoteID"`
	SaleableItems  []*SaleableItem `json:"saleableItems,omitempty" bson:"saleableItems,omitempty"`
	SpecialPricing *SpecialPricing `json:"specialPricing,omitempty" bson:"specialPricing,omitempty"`
	Status         *QuoteStatus    `json:"status,omitempty" bson:"status,omitempty"`
	UpdatedAt      *time.Time      `json:"updatedAt" bson:"updatedAt"`
}

type QuoteStatus struct {
	Active    *bool             `json:"active,omitempty" bson:"active,omitempty"`
	CreatedBy string            `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedBy string            `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	Meta      map[string]string `json:"meta,omitempty" bson:"meta,omitempty"`
	Version   time.Time         `json:"version,omitempty" bson:"version,omitempty"`
}

type SpecialPricing struct {
	IsPolitical               *bool `json:"isPolitical,omitempty" bson:"isPolitical,omitempty"`
	UseExternalProductionCost *bool `json:"useExternalProductionCost,omitempty" bson:"useExternalProductionCost,omitempty"`
}

type SaleableItem struct {
	BuyType                     BuyType  `json:"buyType,omitempty" bson:"buyType,omitempty"`
	EffectiveQuantityGroupTotal float64  `json:"effectiveQuantityGroupTotal,omitempty" bson:"effectiveQuantityGroupTotal,omitempty"`
	ExternalIDs                 []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	// includes
	// salesforce:panelLineItem:a3JVT000001Kh532AC
	ID                  string           `json:"saleableItemID,omitempty" bson:"saleableItemID,omitempty"`
	IsTaxExempt         *bool            `json:"isTaxExempt,omitempty" bson:"isTaxExempt,omitempty"`
	MarketCode          string           `json:"marketCode,omitempty" bson:"marketCode,omitempty"`
	MarketDurationTotal *int             `json:"marketDurationTotal,omitempty" bson:"marketDurationTotal,omitempty"`
	MediaGrouping       string           `json:"mediaGrouping,omitempty" bson:"mediaGrouping,omitempty"`
	PricedItems         []*PricedItem    `json:"pricedItems,omitempty" bson:"pricedItems,omitempty"`
	PricingGuidance     *PricingGuidance `json:"pricingGuidance,omitempty" bson:"pricingGuidance,omitempty"`
	Quantity            int              `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Schedule            Schedule         `json:"schedule,omitempty" bson:"schedule,omitempty"`
}

type Debug struct {
	ForcedProduction *bool                                `json:"forcedProduction,omitempty" bson:"forcedProduction,omitempty"`
	Modifiers        *map[ModifierType][]ModifierSnapshot `json:"modifiers,omitempty" bson:"modifiers,omitempty"`
	Rates            *RateGuidance                        `json:"rates,omitempty" bson:"rates,omitempty"`
	TotalDays        int                                  `json:"totalDays,omitempty" bson:"totalDays,omitempty"`
}

type Cost struct {
	CostPerSqFt *float64 `json:"baseCostPerSqFt,omitempty" bson:"baseCostPerSqFt,omitempty"`
	Tax         float64  `json:"tax" bson:"tax,omitempty"`
	TaxRate     float64  `json:"taxRate,omitempty" bson:"taxRate,omitempty"`
	Value       float64  `json:"value" bson:"value,omitempty"`
}

type Costs struct {
	Install    *Cost `json:"install,omitempty" bson:"install,omitempty"`
	Production *Cost `json:"production,omitempty" bson:"production,omitempty"`
	Repost     *Cost `json:"repost,omitempty" bson:"repost,omitempty"`
}

type Price struct {
	Tax   float64 `json:"tax" bson:"tax,omitempty"`
	Value float64 `json:"value" bson:"value,omitempty"`
}

type PricedItem struct {
	DisplayID   string   `json:"displayID,omitempty" bson:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	// includes
	// salesforce:display:a3IVF000000YNLZ2A4
	InstallTaxRate    float64          `json:"installTaxRate" bson:"installTaxRate,omitempty"`
	PricingGuidance   *PricingGuidance `json:"pricingGuidance,omitempty" bson:"pricingGuidance,omitempty"`
	ProductionTaxRate float64          `json:"productionTaxRate" bson:"productionTaxRate,omitempty"`
	Slots             int              `json:"slots,omitempty" bson:"slots,omitempty"`
}

type Prices struct {
	Base     *Price `json:"base,omitempty" bson:"base,omitempty"`
	Floor    *Price `json:"floor,omitempty" bson:"floor,omitempty"`
	Go       *Price `json:"go,omitempty" bson:"go,omitempty"`
	Min      *Price `json:"min,omitempty" bson:"min,omitempty"`
	Proposed *Price `json:"proposed,omitempty" bson:"proposed,omitempty"`
	RateCard *Price `json:"rateCard,omitempty" bson:"rateCard,omitempty"`
}

type PricingGuidance struct {
	Debug  *Debug  `json:"debug,omitempty" bson:"debug,omitempty"`
	Costs  Costs   `json:"costs,omitempty" bson:"costs,omitempty"`
	Prices Prices  `json:"prices,omitempty" bson:"prices,omitempty"`
	Totals *Totals `json:"totals,omitempty" bson:"totals,omitempty"`
}

type Totals struct {
	Subtotal float64 `json:"subtotal" bson:"subtotal,omitempty"`
	Tax      float64 `json:"tax" bson:"tax,omitempty"`
	TaxRate  float64 `json:"taxRate" bson:"taxRate,omitempty"`
	Total    float64 `json:"total" bson:"total,omitempty"`
}

/* modifier snapshot */
type ModifierCriteria struct {
	BuyType         *BuyType       `json:"buyType,omitempty" bson:"buyType,omitempty"`
	DurationDays    *ModifierRange `json:"durationDays,omitempty" bson:"durationDays,omitempty"`
	EndDate         *time.Time     `json:"endDate,omitempty" bson:"endDate,omitempty"`
	MarketCode      string         `json:"marketCode,omitempty" bson:"marketCode,omitempty"`
	MediaType       *MediaType     `json:"mediaType,omitempty" bson:"mediaType,omitempty"`
	MediaTypeDetail string         `json:"mediaTypeDetail,omitempty" bson:"mediaTypeDetail,omitempty"`
	Quantity        *ModifierRange `json:"quantity,omitempty" bson:"quantity,omitempty"`
	StartDate       *time.Time     `json:"startDate,omitempty" bson:"startDate,omitempty"`
	ValueTier       *ValueTier     `json:"valueTier,omitempty" bson:"valueTier,omitempty"`
}

type ModifierRange struct {
	Max int `json:"max,omitempty" bson:"max,omitempty"`
	Min int `json:"min,omitempty" bson:"min,omitempty"`
}

type ModifierSnapshot struct {
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ModifierCriteria
	ModifierID  string  `json:"modifierID,omitempty" bson:"modifierID,omitempty"`
	OverlapDays int     `json:"overlapDays,omitempty" bson:"overlapDays,omitempty"`
	Percent     float64 `json:"percent" bson:"percent,omitempty"`
	Value       float64 `json:"value,omitempty" bson:"value,omitempty"`
}

/* schedule */
type DateRange struct {
	StartDate time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
}

type Schedule []DateRange

/* rate guidance */
type RateGuidance struct {
	BaseRate      float64        `json:"baseRate,omitempty" bson:"baseRate,omitempty"`
	Overrides     []RateOverride `json:"overrides,omitempty" bson:"overrides,omitempty"`
	PoliticalRate float64        `json:"politicalRate,omitempty" bson:"politicalRate,omitempty"`
	RateCard      float64        `json:"rateCard,omitempty" bson:"rateCard,omitempty"`
}

type RateOverride struct {
	BaseRate          *float64       `json:"baseRate,omitempty" bson:"baseRate,omitempty"`
	EndDate           time.Time      `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExcludedModifiers []ModifierType `json:"excludedModifiers,omitempty" bson:"excludedModifiers,omitempty"`
	ExternalIDs       []string       `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	RateCard          *float64       `json:"rateCard,omitempty" bson:"rateCard,omitempty"`
	StartDate         time.Time      `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

/* enums */
type BuyType string

const (
	ImpressionsFlex         BuyType = "ImpressionsFlexible"
	ImpressionsFixed        BuyType = "ImpressionsFixed"
	DisplayFixed            BuyType = "DisplayFixed"
	DisplayFlex             BuyType = "DisplayFlexible"
	Network                 BuyType = "Network"
	NonTrad                 BuyType = "Non-Traditional"
	QuantityFixed           BuyType = "QuantityFixed"
	QuantityFlex            BuyType = "QuantityFlexible"
	QuantityFullMarketFixed BuyType = "QuantityFullMarketFixed"
	QuantityFullMarketFlex  BuyType = "QuantityFullMarketFlexible"
)

type MediaType string

const (
	AirportDisplays       MediaType = "Airport Displays"
	Backlit               MediaType = "Backlit"
	BacklitKiosk          MediaType = "Backlit Kiosk"
	Banners               MediaType = "Banners"
	Bench                 MediaType = "Bench"
	Bulletin              MediaType = "Bulletin"
	BusExterior           MediaType = "Bus Exterior"
	BusInterior           MediaType = "Bus Interior"
	BusKing               MediaType = "Bus King"
	BusQueen              MediaType = "Bus Queen"
	BusTail               MediaType = "Bus Tail"
	BusWrap               MediaType = "Bus Wrap"
	DigitalBulletin       MediaType = "Digital Bulletin"
	DigitalDisplay        MediaType = "Digital Display"
	DigitalPremierePanel  MediaType = "Digital Premiere Panel"
	DigitalSpectacular    MediaType = "Digital Spectacular"
	DigitalTransit        MediaType = "Digital Transit"
	DigitalTransitShelter MediaType = "Digital Transit Shelter"
	DigitalUrbanPanel     MediaType = "Digital Urban Panel"
	ExhibitsMisc          MediaType = "Exhibits/Misc."
	JuniorPoster          MediaType = "Junior Poster"
	Kiosk                 MediaType = "Kiosk"
	Malls                 MediaType = "Malls"
	MobileBillboards      MediaType = "Mobile Billboards"
	Newsrack              MediaType = "Newsrack"
	Other                 MediaType = "Other"
	OtherVenues           MediaType = "Other Venues"
	Poster                MediaType = "Poster"
	PremiereJuniorPoster  MediaType = "Premiere Junior Poster"
	PremierePanel         MediaType = "Premiere Panel"
	PremiereSquare        MediaType = "Premiere Square"
	Spectacular           MediaType = "Spectacular"
	TrainInterior         MediaType = "Train Interior"
	TransitDisplays       MediaType = "Transit Displays"
	TransitShelter        MediaType = "Transit Shelter"
	UnknownMedia          MediaType = ""
	Wallscape             MediaType = "Wallscape"
	WrapsClings           MediaType = "Wraps/Clings"
)

type ValueTier string

const (
	Economy      ValueTier = "Economy"
	Premium      ValueTier = "Premium"
	Standard     ValueTier = "Standard"
	StandardPlus ValueTier = "Standard+"
	Target       ValueTier = "Target"
	UnknownValue ValueTier = ""
	Value        ValueTier = "Value"
)

type ModifierType string

const (
	DynamicMod             ModifierType = "dynamic"
	EventMod               ModifierType = "event"
	FlexibleMod            ModifierType = "flexible"
	FullMarketMod          ModifierType = "fullMarket"
	FullMarketPoliticalMod ModifierType = "fullMarketPolitical"
	PoliticalRateMod       ModifierType = "political"
	ProductionCostMod      ModifierType = "productionCost"
	RateMod                ModifierType = "rate"
	TraditionalMod         ModifierType = "traditional"
)
