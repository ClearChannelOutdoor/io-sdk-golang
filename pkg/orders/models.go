package orders

import "time"

type AdditionalCost struct {
	ContractValue *float64           `json:"contractValue,omitempty" bson:"contractValue,omitempty"`
	Name          AdditionalCostName `json:"name,omitempty" bson:"name,omitempty"`
}

type Order struct {
	AdditionalCosts []*AdditionalCost `json:"additionalCosts,omitempty" bson:"additionalCosts,omitempty"`

	AdvertisedProductCode   string         `json:"advertisedProductCode,omitempty"`
	Advertiser              CustomerEntity `json:"advertiser,omitempty"`
	Buyer                   CustomerEntity `json:"buyer,omitempty"`
	Canceled                *bool          `json:"canceled,omitempty"`
	CancellationTerms       string         `json:"cancellationTerms,omitempty"`
	CreatedAt               time.Time      `json:"createdAt,omitempty"`
	CreativeChangeAgreement *bool          `json:"creativeChangeAgreement,omitempty"`
	EndDate                 time.Time      `json:"endDate,omitempty"`
	ExternalIDs             []string       `json:"externalIDs,omitempty"`
	ID                      string         `json:"orderID,omitempty"`
	Markets                 []OrderMarket  `json:"markets,omitempty"`
	Name                    string         `json:"name,omitempty"`
	Number                  string         `json:"number,omitempty"`
	Psa                     *bool          `json:"psa,omitempty"`
	Seller                  *OrderSeller   `json:"seller,omitempty"`
	Source                  OrderSource    `json:"source,omitempty"`
	StartDate               time.Time      `json:"startDate,omitempty"`
	UpdatedAt               time.Time      `json:"updatedAt,omitempty"`
}

type CustomerEntity struct {
	ID           string `json:"customerID,omitempty"`
	IsAdvertiser bool   `json:"isAdvertiser,omitempty"`
	Number       string `json:"number,omitempty"`
}

type Employee struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Number      string   `json:"number,omitempty"`
}

type OrderMarket struct {
	Canceled           bool               `json:"canceled,omitempty"`
	CancellationReason string             `json:"cancellationReason,omitempty"`
	ExternalIDs        []string           `json:"externalIDs,omitempty"`
	MarketCode         string             `json:"marketCode,omitempty"`
	MarketID           string             `json:"marketID,omitempty"`
	SubmissionCategory SubmissionCategory `json:"submissionCategory,omitempty"`
}

type OrderSeller struct {
	LeadAccountExec         Employee `json:"leadAccountExec,omitempty"`
	SellingBusinessUnitCode string   `json:"sellingBusinessUnitCode,omitempty"`
	SupportingAccountExec   Employee `json:"supportingAccountExec,omitempty"`
}
