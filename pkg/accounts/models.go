package accounts

import "time"

type Account struct {
	ID          string           `json:"accountID"`
	Advertiser  *AccountCustomer `json:"advertiser,omitempty" bson:"-"`
	BillTo      *AccountCustomer `json:"billTo,omitempty" bson:"-"`
	Buyer       *AccountCustomer `json:"buyer,omitempty" bson:"-"`
	CreatedAt   time.Time        `json:"createdAt"`
	ExternalIDs []string         `json:"externalIDs,omitempty"`
	ThirdParty  *AccountCustomer `json:"thirdParty,omitempty" bson:"-"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type AccountCustomer struct {
	ID          string   `json:"customerID"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Name        string   `json:"name"`
}
