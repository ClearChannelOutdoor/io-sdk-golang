package customers

import "time"

type Category string

const (
	Local    Category = "Local"
	National Category = "National"
)

type Classification string

const (
	DoNotCall     Classification = "Do Not Call"
	Key           Classification = "Key"
	OutOfBusiness Classification = "Out Of Business"
	Prospect      Classification = "Prospect"
	Secondary     Classification = "Secondary"
	Target        Classification = "Target"
)

type Credit struct {
	ApprovedAt time.Time `json:"approvedAt,omitempty"`
	Limit      int       `json:"limit,omitempty"`
	Notes      string    `json:"notes,omitempty"`
	Status     Status    `json:"status,omitempty"`
	Terms      Terms     `json:"terms,omitempty"`
}

type Customer struct {
	ID           string    `json:"customerID"`
	Code         string    `json:"code,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	Credit       *Credit   `json:"credit,omitempty"`
	ExternalIDs  []string  `json:"externalIDs,omitempty"`
	IsAdvertiser bool      `json:"isAdvertiser"`
	Markets      []Market  `json:"markets"`
	NAICS        *NAICS    `json:"naics,omitempty"`
	Name         string    `json:"name"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Market struct {
	ID             string         `json:"marketID,omitempty"`
	Category       Category       `json:"category,omitempty"`
	Classification Classification `json:"classification,omitempty"`
	CreatedAt      time.Time      `json:"createdAt"`
	Credit         *Credit        `json:"credit,omitempty"`
	ExternalIDs    []string       `json:"externalIDs,omitempty"`
	Name           string         `json:"name"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

type NAICS struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type Status string

const (
	Active      Status = "Active"
	Deactivated Status = "Deactivated"
	Deleted     Status = "Deleted"
	Merged      Status = "Merged"
	Temporary   Status = "Temporary"
)

type Terms string

const (
	DUR    Terms = "DUR"
	NET30  Terms = "NET30"
	NET45  Terms = "NET45"
	NET60  Terms = "NET60"
	NET90  Terms = "NET90"
	NET120 Terms = "NET120"
	PPIF   Terms = "PPIF"
	PPIP   Terms = "PPIP"
	PRE    Terms = "PRE"
)
