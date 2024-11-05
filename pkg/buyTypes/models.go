package buyTypes

import (
	"time"
)

type BuyType struct {
	Attributes  *BuyTypeAttributes `json:"attributes,omitempty" bson:"attributes,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Codes       *BuyTypeCodes      `json:"codes,omitempty" bson:"codes,omitempty"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	ExternalIDs []string           `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string             `json:"buyTypeID,omitempty" bson:"buyTypeID,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type BuyTypeAttributes struct {
	ContractSettings string `json:"contractSettings,omitempty" bson:"contractSettings,omitempty"`
	Moveable         *bool  `json:"moveable,omitempty" bson:"moveable,omitempty"`
	NotBillable      *bool  `json:"notBillable,omitempty" bson:"notBillable,omitempty"`
	NotChargeable    *bool  `json:"notChargeable,omitempty" bson:"notChargeable,omitempty"`
	PSA              *bool  `json:"psa,omitempty" bson:"psa,omitempty"`
	Suspendable      *bool  `json:"suspendable,omitempty" bson:"suspendable,omitempty"`
}

type BuyTypeCodes struct {
	DetailCode         string `json:"detailCode,omitempty" bson:"detailCode,omitempty"`
	DetailDescription  string `json:"detailDescription,omitempty" bson:"detailDescription,omitempty"`
	SegmentCode        string `json:"segmentCode,omitempty" bson:"segmentCode,omitempty"`
	SegmentDescription string `json:"segmentDescription,omitempty" bson:"segmentDescription,omitempty"`
}
