package restrictions

import "time"

type CCOProductCode struct {
	Code     string `json:"code,omitempty" bson:"code,omitempty"`
	FullCode string `json:"fullCode,omitempty" bson:"fullCode,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
}

type IABTaxonomy struct {
	V1 string `json:"v1,omitempty" bson:"v1,omitempty"`
	V2 string `json:"v2,omitempty" bson:"v2,omitempty"`
	V3 string `json:"v3,omitempty" bson:"v3,omitempty"`
}

type Restriction struct {
	Allowed     *bool                `json:"allowed,omitempty" bson:"allowed,omitempty"`
	CreatedAt   *time.Time           `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Display     *RestrictionDisplay  `json:"display,omitempty" bson:"display,omitempty"`
	EndAt       *time.Time           `json:"endAt,omitempty" bson:"endAt,omitempty"`
	ExternalIDs []string             `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string               `json:"restrictionID" bson:"restrictionID"`
	Notes       *string              `json:"notes,omitempty" bson:"notes,omitempty"`
	StartAt     *time.Time           `json:"startAt,omitempty" bson:"startAt,omitempty"`
	Taxonomy    *RestrictionTaxonomy `json:"taxonomy,omitempty" bson:"taxonomy,omitempty"`
	UpdatedAt   *time.Time           `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type RestrictionDisplay struct {
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string   `json:"displayID,omitempty" bson:"displayID,omitempty"`
}

type RestrictionTaxonomy struct {
	CCO *CCOProductCode `json:"cco,omitempty" bson:"cco,omitempty"`
	IAB *IABTaxonomy    `json:"iab,omitempty" bson:"iab,omitempty"`
}
