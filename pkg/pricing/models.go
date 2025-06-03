package pricing

import (
	"time"
)

type Quote struct {
	CreatedAt   *time.Time `json:"createdAt"`
	ExternalIDs []string   `json:"externalIDs,omitempty"`
	// includes
	// salesforce:opportunity:a3HVT000001HYBB2A4
	// salesforce:plan:a3JVT000001Kh532AC
	ID             string          `json:"quoteID"`
	SaleableItems  []*SaleableItem `json:"saleableItems,omitempty"`
	SpecialPricing *SpecialPricing `json:"specialPricing,omitempty"`
	Status         *QuoteStatus    `json:"status,omitempty"`
	UpdatedAt      *time.Time      `json:"updatedAt"`
}

type QuoteStatus struct {
	Active    *bool             `json:"active,omitempty"`
	CreatedBy string            `json:"createdBy,omitempty"`
	UpdatedBy string            `json:"updatedBy,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
	Version   time.Time         `json:"version,omitempty"`
}

type SpecialPricing struct {
	IsPolitical               *bool `json:"isPolitical,omitempty"`
	UseExternalProductionCost *bool `json:"useExternalProductionCost,omitempty"`
}

type SaleableItem struct {
	BuyType                     BuyType  `json:"buyType,omitempty"`
	EffectiveQuantityGroupTotal float64  `json:"effectiveQuantityGroupTotal,omitempty"`
	ExternalIDs                 []string `json:"externalIDs,omitempty"`
	// includes
	// salesforce:panelLineItem:a3JVT000001Kh532AC
	ID                  string           `json:"saleableItemID,omitempty"`
	IsTaxExempt         *bool            `json:"isTaxExempt,omitempty"`
	MarketCode          string           `json:"marketCode,omitempty"`
	MarketDurationTotal *int             `json:"marketDurationTotal,omitempty"`
	MediaGrouping       string           `json:"mediaGrouping,omitempty"`
	PricedItems         []*PricedItem    `json:"pricedItems,omitempty"`
	PricingGuidance     *PricingGuidance `json:"pricingGuidance,omitempty"`
	Quantity            int              `json:"quantity,omitempty"`
	Schedule            Schedule         `json:"schedule,omitempty"`
}

type Debug struct {
	ForcedProduction *bool                                `json:"forcedProduction,omitempty"`
	Modifiers        *map[ModifierType][]ModifierSnapshot `json:"modifiers,omitempty"`
	Rates            *RateGuidance                        `json:"rates,omitempty"`
	TotalDays        int                                  `json:"totalDays,omitempty"`
}

type Cost struct {
	CostPerSqFt *float64 `json:"baseCostPerSqFt,omitempty"`
	Tax         float64  `json:"tax"`
	TaxRate     float64  `json:"taxRate,omitempty"`
	Value       float64  `json:"value"`
}

type Costs struct {
	Install    *Cost `json:"install,omitempty"`
	Production *Cost `json:"production,omitempty"`
	Repost     *Cost `json:"repost,omitempty"`
}

type Price struct {
	Tax   float64 `json:"tax"`
	Value float64 `json:"value"`
}

type PricedItem struct {
	DisplayID   string   `json:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	// includes
	// salesforce:display:a3IVF000000YNLZ2A4
	InstallTaxRate    float64          `json:"installTaxRate"`
	PricingGuidance   *PricingGuidance `json:"pricingGuidance,omitempty"`
	ProductionTaxRate float64          `json:"productionTaxRate"`
	Slots             int              `json:"slots,omitempty"`
}

type Prices struct {
	Base     *Price `json:"base,omitempty"`
	Floor    *Price `json:"floor,omitempty"`
	Go       *Price `json:"go,omitempty"`
	Min      *Price `json:"min,omitempty"`
	Proposed *Price `json:"proposed,omitempty"`
	RateCard *Price `json:"rateCard,omitempty"`
}

type PricingGuidance struct {
	Debug  *Debug  `json:"debug,omitempty"`
	Costs  Costs   `json:"costs,omitempty"`
	Prices Prices  `json:"prices,omitempty"`
	Totals *Totals `json:"totals,omitempty"`
}

type Totals struct {
	Subtotal float64 `json:"subtotal"`
	Tax      float64 `json:"tax"`
	TaxRate  float64 `json:"taxRate"`
	Total    float64 `json:"total"`
}

/* modifier snapshot */
type ModifierCriteria struct {
	BuyType         *BuyType       `json:"buyType,omitempty"`
	DurationDays    *ModifierRange `json:"durationDays,omitempty"`
	EndDate         *time.Time     `json:"endDate,omitempty"`
	MarketCode      string         `json:"marketCode,omitempty"`
	MediaType       *MediaType     `json:"mediaType,omitempty"`
	MediaTypeDetail string         `json:"mediaTypeDetail,omitempty"`
	Quantity        *ModifierRange `json:"quantity,omitempty"`
	StartDate       *time.Time     `json:"startDate,omitempty"`
	ValueTier       *ValueTier     `json:"valueTier,omitempty"`
}

type ModifierRange struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}

type ModifierSnapshot struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ModifierCriteria
	ModifierID  string  `json:"modifierID,omitempty"`
	OverlapDays int     `json:"overlapDays,omitempty"`
	Percent     float64 `json:"percent"`
	Value       float64 `json:"value,omitempty"`
}

/* schedule */
type DateRange struct {
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
}

type Schedule []DateRange

/* rate guidance */
type RateGuidance struct {
	BaseRate      float64        `json:"baseRate,omitempty"`
	Overrides     []RateOverride `json:"overrides,omitempty"`
	PoliticalRate float64        `json:"politicalRate,omitempty"`
	RateCard      float64        `json:"rateCard,omitempty"`
}

type RateOverride struct {
	BaseRate          *float64       `json:"baseRate,omitempty"`
	EndDate           time.Time      `json:"endDate,omitempty"`
	ExcludedModifiers []ModifierType `json:"excludedModifiers,omitempty"`
	ExternalIDs       []string       `json:"externalIDs,omitempty"`
	RateCard          *float64       `json:"rateCard,omitempty"`
	StartDate         time.Time      `json:"startDate,omitempty"`
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
