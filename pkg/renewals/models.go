package renewals

import "time"

type Relationship struct {
	CreatedAt *time.Time        `json:"createdAt,omitempty"`
	ID        string            `json:"relationshipID,omitempty"`
	Origin    *SystemIdentifier `json:"origin,omitempty"`
	Renewal   *SystemIdentifier `json:"renewal,omitempty"`
	UpdatedAt *time.Time        `json:"updatedAt,omitempty"`
}

type SystemIdentifier struct {
	Entity      string   `json:"entity,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	ID          string   `json:"id,omitempty"`
	Source      string   `json:"source,omitempty"`
}
