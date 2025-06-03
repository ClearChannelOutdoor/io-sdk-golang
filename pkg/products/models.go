package products

import "time"

type Media struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty"`
	ProductCode string   `json:"productCode"`
	Type        string   `json:"type"`
	TypeCode    string   `json:"typeCode"`
}

type Product struct {
	ID          string    `json:"productID"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty"`
	IsDigital   bool      `json:"isDigital"`
	Media       []Media   `json:"media,omitempty"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
