package orders

import "time"

/* Orders */
type Order struct {
	AdvertisedProductCode   string         `json:"advertisedProductCode,omitempty" bson:"advertisedProductCode,omitempty"`
	Advertiser              CustomerEntity `json:"advertiser,omitempty" bson:"advertiser,omitempty"`
	Buyer                   CustomerEntity `json:"buyer,omitempty" bson:"buyer,omitempty"`
	Canceled                *bool          `json:"canceled" bson:"canceled"`
	CancellationTerms       string         `json:"cancellationTerms,omitempty" bson:"cancellationTerms,omitempty"`
	CancellationReason      string         `json:"cancellationReason,omitempty" bson:"cancellationReason,omitempty" `
	CreatedAt               time.Time      `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	CreativeChangeAgreement *bool          `json:"creativeChangeAgreement,omitempty" bson:"creativeChangeAgreement,omitempty"`
	EndDate                 time.Time      `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ID                      string         `json:"orderID,omitempty" bson:"orderID,omitempty"`
	Markets                 []OrderMarket  `json:"markets,omitempty" bson:"markets,omitempty" `
	Number                  string         `json:"number,omitempty" bson:"number,omitempty"`
	Name                    string         `json:"name,omitempty" bson:"name,omitempty"`
	Psa                     *bool          `json:"psa,omitempty" bson:"psa,omitempty"`
	Seller                  *OrderSeller   `json:"seller,omitempty" bson:"seller,omitempty"`
	Source                  OrderSource    `json:"source,omitempty" bson:"source,omitempty"`
	StartDate               time.Time      `json:"startDate,omitempty" bson:"startDate,omitempty"`
	UpdatedAt               time.Time      `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`

	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type OrderMarket struct {
	Canceled           bool               `json:"canceled,omitempty" bson:"canceled,omitempty"`
	MarketID           string             `json:"marketID,omitempty" bson:"marketID,omitempty"`
	MarketCode         string             `json:"marketCode,omitempty" bson:"marketCode,omitempty"`
	SubmissionCategory SubmissionCategory `json:"submissionCategory,omitempty" bson:"submissionCategory,omitempty"`
	ExternalIDs        []string           `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}
type Employee struct {
	Number      string   `json:"number,omitempty" bson:"number,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type OrderSeller struct {
	LeadAccountExec         Employee `json:"leadAccountExec,omitempty" bson:"leadAccountExec,omitempty"`
	SellingBusinessUnitCode string   `json:"sellingBusinessUnitCode,omitempty" bson:"sellingBusinessUnitCode,omitempty"`
	SupportingAccountExec   Employee `json:"supportingAccountExec,omitempty" bson:"supportingAccountExec,omitempty"`
}

type CustomerEntity struct {
	ID           string `json:"customerID" bson:"customerID"`
	Number       string `json:"number,omitempty" bson:"number,omitempty"`
	IsAdvertiser bool   `json:"isAdvertiser,omitempty" bson:"isAdvertiser,omitempty"`
}
