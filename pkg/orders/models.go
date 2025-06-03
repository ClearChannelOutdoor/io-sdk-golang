package orders

import "time"

/* Orders */
type Order struct {
	AdvertisedProductCode   string         `json:"advertisedProductCode,omitempty"`
	Advertiser              CustomerEntity `json:"advertiser,omitempty"`
	Buyer                   CustomerEntity `json:"buyer,omitempty"`
	Canceled                *bool          `json:"canceled"`
	CancellationTerms       string         `json:"cancellationTerms,omitempty"`
	CancellationReason      string         `json:"cancellationReason,omitempty" `
	CreatedAt               time.Time      `json:"createdAt,omitempty"`
	CreativeChangeAgreement *bool          `json:"creativeChangeAgreement,omitempty"`
	EndDate                 time.Time      `json:"endDate,omitempty"`
	ID                      string         `json:"orderID,omitempty"`
	Markets                 []OrderMarket  `json:"markets,omitempty" `
	Number                  string         `json:"number,omitempty"`
	Name                    string         `json:"name,omitempty"`
	Psa                     *bool          `json:"psa,omitempty"`
	Seller                  *OrderSeller   `json:"seller,omitempty"`
	Source                  OrderSource    `json:"source,omitempty"`
	StartDate               time.Time      `json:"startDate,omitempty"`
	UpdatedAt               time.Time      `json:"updatedAt,omitempty"`

	ExternalIDs []string `json:"externalIDs,omitempty"`
}

type OrderMarket struct {
	Canceled           bool               `json:"canceled,omitempty"`
	MarketID           string             `json:"marketID,omitempty"`
	MarketCode         string             `json:"marketCode,omitempty"`
	SubmissionCategory SubmissionCategory `json:"submissionCategory,omitempty"`
	ExternalIDs        []string           `json:"externalIDs,omitempty"`
}
type Employee struct {
	Number      string   `json:"number,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
}

type OrderSeller struct {
	LeadAccountExec         Employee `json:"leadAccountExec,omitempty"`
	SellingBusinessUnitCode string   `json:"sellingBusinessUnitCode,omitempty"`
	SupportingAccountExec   Employee `json:"supportingAccountExec,omitempty"`
}

type CustomerEntity struct {
	ID           string `json:"customerID"`
	Number       string `json:"number,omitempty"`
	IsAdvertiser bool   `json:"isAdvertiser,omitempty"`
}
