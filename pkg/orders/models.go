package orders

import "time"

/* Orders */
type Order struct {
	Account                 *OrderAccount `json:"account,omitempty" bson:"account,omitempty"`
	AdvertisedProductCode   string        `json:"advertisedProductCode,omitempty" bson:"advertisedProductCode,omitempty"`
	BookingNotes            string        `json:"bookingNotes,omitempty" bson:"bookingNotes,omitempty"`
	Canceled                bool          `json:"canceled" bson:"canceled"`
	CancellationTerms       string        `json:"cancellationTerms,omitempty" bson:"cancellationTerms,omitempty"`
	CancellationReason      string        `json:"cancellationReason,omitempty" bson:"cancellationReason,omitempty" `
	CreatedAt               time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	CreativeChangeAgreement bool          `json:"creativeChangeAgreement,omitempty" bson:"creativeChangeAgreement,omitempty"`
	EndDate                 *time.Time    `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ID                      string        `json:"id,omitempty" bson:"id,omitempty"`
	Markets                 []OrderMarket `json:"markets,omitempty" bson:"markets,omitempty" `
	Number                  string        `json:"number,omitempty" bson:"number,omitempty"`
	Name                    string        `json:"name,omitempty" bson:"name,omitempty"`
	Psa                     bool          `json:"psa,omitempty" bson:"psa,omitempty"`
	Seller                  *OrderSeller  `json:"seller,omitempty" bson:"seller,omitempty"`
	Source                  OrderSource   `json:"source,omitempty" bson:"source,omitempty"`
	StartDate               *time.Time    `json:"startDate,omitempty" bson:"startDate,omitempty"`
	UpdatedAt               time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`

	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type OrderMarket struct {
	MarketID           string             `json:"marketID,omitempty" bson:"marketID,omitempty"`
	MarketCode         string             `json:"marketCode,omitempty" bson:"marketCode,omitempty"`
	SubmissionCategory SubmissionCategory `json:"submissionCategory,omitempty" bson:"submissionCategory,omitempty"`
	ExternalIDs        []string           `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type Employee struct {
	FullName    string   `json:"fullName,omitempty" bson:"fullName,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type OrderAccount struct {
	AccountID             string         `json:"accountID" bson:"accountID"`
	Advertiser            CustomerEntity `json:"advertiser,omitempty" bson:"advertiser,omitempty"`
	Buyer                 CustomerEntity `json:"buyer,omitempty" bson:"buyer,omitempty"`
	DefaultBillToCustomer CustomerEntity `json:"defaultBillToCustomer,omitempty" bson:"defaultBillToCustomer,omitempty"`
	ExternalIDs           []string       `json:"externalIDs"`
}

type OrderSeller struct {
	LeadAccountExec         Employee `json:"leadAccountExec,omitempty" bson:"leadAccountExec,omitempty"`
	SellingBusinessUnitCode string   `json:"sellingBusinessUnitCode,omitempty" bson:"sellingBusinessUnitCode,omitempty"`
	SupportingAccountExec   Employee `json:"supportingAccountExec,omitempty" bson:"supportingAccountExec,omitempty"`
}

type CustomerEntity struct {
	CustomerID  string   `json:"customerID" bson:"customerID"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
	Number      string   `json:"number,omitempty" bson:"number,omitempty"`
	Type        string   `json:"type,omitempty" bson:"type,omitempty"`
}

/* OrderLines */
type OrderLineStatus string

const (
	Draft     OrderLineStatus = "Draft"
	Reserved  OrderLineStatus = "Reserved"
	Booked    OrderLineStatus = "Booked"
	Cancelled OrderLineStatus = "Cancelled"
)

type OrderLine struct {
	CreatedAt          time.Time       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Description        string          `json:"description,omitempty" bson:"description,omitempty"`
	EndDate            *time.Time      `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExternalIDs        []string        `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID                 string          `json:"orderLineID,omitempty" bson:"orderLineID,omitempty"`
	MarketID           string          `json:"marketID,omitempty" bson:"marketID,omitempty"`
	MediaProduct       *MediaProduct   `json:"mediaProduct,omitempty" bson:"mediaProduct,omitempty"`
	OrderID            string          `json:"orderID,omitempty" bson:"orderID,omitempty"`
	Quantity           *int            `json:"quantity,omitempty" bson:"quantity,omitempty"`
	StartDate          *time.Time      `json:"startDate,omitempty" bson:"startDate,omitempty"`
	Status             OrderLineStatus `json:"status,omitempty" bson:"status,omitempty"`
	TargetRatingPoints *int            `json:"targetRatingPoints,omitempty" bson:"targetRatingPoints,omitempty"`
	Type               *Type           `json:"type,omitempty" bson:"type,omitempty"`
	UpdatedAt          time.Time       `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type MediaProduct struct {
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ProductCode string   `json:"productCode,omitempty" bson:"productCode,omitempty"`
	TypeCode    string   `json:"typeCode,omitempty" bson:"typeCode,omitempty"`
}

type Type struct {
	DetailCode  string   `json:"detailCode,omitempty" bson:"detailCode,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	NoBill      bool     `json:"noBill,omitempty" bson:"noBill,omitempty"`
	SegmentCode string   `json:"segmentCode,omitempty" bson:"segmentCode,omitempty"`
}
