package restrictions

import "time"

type CCOProductCode struct {
	Code     string `json:"code,omitempty"`
	FullCode string `json:"fullCode,omitempty"`
	Name     string `json:"name,omitempty"`
}

type IABTaxonomy struct {
	V1 string `json:"v1,omitempty"`
	V2 string `json:"v2,omitempty"`
	V3 string `json:"v3,omitempty"`
}

type Restriction struct {
	Allowed     *bool                `json:"allowed,omitempty"`
	CreatedAt   *time.Time           `json:"createdAt,omitempty"`
	Display     *RestrictionDisplay  `json:"display,omitempty"`
	EndAt       *time.Time           `json:"endAt,omitempty"`
	ExternalIDs []string             `json:"externalIDs,omitempty"`
	ID          string               `json:"restrictionID"`
	Notes       *string              `json:"notes,omitempty"`
	StartAt     *time.Time           `json:"startAt,omitempty"`
	Taxonomy    *RestrictionTaxonomy `json:"taxonomy,omitempty"`
	UpdatedAt   *time.Time           `json:"updatedAt,omitempty"`
}

type RestrictionDisplay struct {
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"displayID,omitempty"`
}

type RestrictionTaxonomy struct {
	CCO *CCOProductCode `json:"cco,omitempty"`
	IAB *IABTaxonomy    `json:"iab,omitempty"`
}
