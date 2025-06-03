package accounts

import "time"

type Account struct {
	ID          string           `json:"accountID"`
	Advertiser  *AccountCustomer `json:"advertiser,omitempty"`
	BillTo      *AccountCustomer `json:"billTo,omitempty"`
	Buyer       *AccountCustomer `json:"buyer,omitempty"`
	CreatedAt   time.Time        `json:"createdAt"`
	ExternalIDs []string         `json:"externalIDs,omitempty"`
	ThirdParty  *AccountCustomer `json:"thirdParty,omitempty"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type AccountCustomer struct {
	ID          string   `json:"customerID"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty"`
	Number      string   `json:"number,omitempty"`
}
