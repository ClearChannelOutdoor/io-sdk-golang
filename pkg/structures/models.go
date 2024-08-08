package structures

import (
	"time"
)

type ExternalID string

func (eid ExternalID) IsEqual(id ExternalID) bool {
	return eid == id
}

type Structure struct {
	ID          string       `json:"structureID" bson:"structureID"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty" bson:"createdAt"`
	ExternalIDs []ExternalID `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Site        *Site        `json:"site,omitempty" bson:"site,omitempty"`
	Type        string       `json:"type,omitempty" bson:"type,omitempty"`
	UpdatedAt   *time.Time   `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type Site struct {
	BusinessUnit string       `json:"businessUnit,omitempty" bson:"businessUnit,omitempty"`
	City         string       `json:"city,omitempty" bson:"city,omitempty"`
	Country      string       `json:"country,omitempty" bson:"country,omitempty"`
	ExternalIDs  []ExternalID `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	State        string       `json:"state,omitempty" bson:"state,omitempty"`
	Street       string       `json:"street,omitempty" bson:"street,omitempty"`
	Title        string       `json:"title,omitempty" bson:"title,omitempty"`
	Zip          string       `json:"zip,omitempty" bson:"zip,omitempty"`
}
