package orders

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

/* Orders */
type Order struct {
	Account                 *OrderAccount `json:"account,omitempty" bson:"account,omitempty"`
	AdvertisedProductCode   string        `json:"advertisedProductCode,omitempty" bson:"advertisedProductCode,omitempty"`
	BookingNotes            string        `json:"bookingNotes,omitempty" bson:"bookingNotes,omitempty"`
	Canceled                *bool         `json:"canceled" bson:"canceled"`
	CancellationTerms       string        `json:"cancellationTerms,omitempty" bson:"cancellationTerms,omitempty"`
	CancellationReason      string        `json:"cancellationReason,omitempty" bson:"cancellationReason,omitempty" `
	CreatedAt               time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	CreativeChangeAgreement *bool         `json:"creativeChangeAgreement,omitempty" bson:"creativeChangeAgreement,omitempty"`
	EndDate                 *time.Time    `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ID                      string        `json:"id,omitempty" bson:"id,omitempty"`
	Markets                 []OrderMarket `json:"markets,omitempty" bson:"markets,omitempty" `
	Number                  string        `json:"number,omitempty" bson:"number,omitempty"`
	Name                    string        `json:"name,omitempty" bson:"name,omitempty"`
	Psa                     *bool         `json:"psa,omitempty" bson:"psa,omitempty"`
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

func (o Order) MarshalZerologObject(e *zerolog.Event) {
	var start, end time.Time
	var canceled, creativeChange, psa string
	if o.StartDate != nil {
		start = *o.StartDate
	}
	if o.EndDate != nil {
		end = *o.EndDate
	}
	if o.Canceled != nil {
		canceled = fmt.Sprintf("%v", *o.Canceled)
	}
	if o.CreativeChangeAgreement != nil {
		creativeChange = fmt.Sprintf("%v", *o.CreativeChangeAgreement)
	}
	if o.Psa != nil {
		psa = fmt.Sprintf("%v", *o.Psa)
	}

	e.
		Str("model", "Order").
		Str("orderID", o.ID).
		Str("advertisedProductCode", o.AdvertisedProductCode).
		Str("canceled", canceled).
		Str("cancellationReason", o.CancellationReason).
		Str("cancellationTerms", o.CancellationTerms).
		Str("creativeChangeAgreement", creativeChange).
		Time("endDate", end).
		Time("startDate", start).
		Str("name", o.Name).
		Str("number", o.Number).
		Str("psa", psa).
		Str("source", string(o.Source))
}
