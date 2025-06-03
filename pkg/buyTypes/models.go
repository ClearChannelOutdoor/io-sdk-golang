package buyTypes

import (
	"time"
)

type BuyType struct {
	Attributes  *BuyTypeAttributes `json:"attributes,omitempty"`
	Category    string             `json:"category,omitempty"`
	Codes       *BuyTypeCodes      `json:"codes,omitempty"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty"`
	Description string             `json:"description,omitempty"`
	ExternalIDs []string           `json:"externalIDs,omitempty"`
	ID          string             `json:"buyTypeID,omitempty"`
	Name        string             `json:"name,omitempty"`
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty"`
}

type BuyTypeAttributes struct {
	ContractSettings string `json:"contractSettings,omitempty"`
	Moveable         *bool  `json:"moveable,omitempty"`
	NotBillable      *bool  `json:"notBillable,omitempty"`
	NotChargeable    *bool  `json:"notChargeable,omitempty"`
	PSA              *bool  `json:"psa,omitempty"`
	Suspendable      *bool  `json:"suspendable,omitempty"`
}

type BuyTypeCodes struct {
	DetailCode         string `json:"detailCode,omitempty"`
	DetailDescription  string `json:"detailDescription,omitempty"`
	SegmentCode        string `json:"segmentCode,omitempty"`
	SegmentDescription string `json:"segmentDescription,omitempty"`
}
