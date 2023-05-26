package products

import "time"

type Media struct {
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
	ProductCode string   `json:"productCode" bson:"productCode"`
	Type        string   `json:"type" bson:"type"`
	TypeCode    string   `json:"typeCode" bson:"typeCode"`
}

type Product struct {
	ID          string    `json:"productID" bson:"productID"`
	Category    string    `json:"category" bson:"category"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	IsDigital   bool      `json:"isDigital" bson:"isDigital"`
	Media       []Media   `json:"media,omitempty" bson:"media,omitempty"`
	Name        string    `json:"name" bson:"name"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
