package renewals

import "time"

type Relationship struct {
	CreatedAt *time.Time        `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ID        string            `json:"relationshipID,omitempty" bson:"relationshipID,omitempty"`
	Origin    *SystemIdentifier `json:"origin,omitempty" bson:"origin,omitempty"`
	Renewal   *SystemIdentifier `json:"renewal,omitempty" bson:"renewal,omitempty"`
	UpdatedAt *time.Time        `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SystemIdentifier struct {
	Entity      string   `json:"entity,omitempty" bson:"entity,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string   `json:"id,omitempty" bson:"id,omitempty"`
	Source      string   `json:"source,omitempty" bson:"source,omitempty"`
}
