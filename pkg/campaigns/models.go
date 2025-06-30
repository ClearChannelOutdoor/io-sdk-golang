package campaigns

import "time"

/* Campaign */
type Audience struct {
	Demographics []AudienceDemographics `json:"demographics,omitempty"`
}

type AudienceDemographics struct {
	Age           string `json:"age,omitempty"`
	Ethnicity     string `json:"ethnicity,omitempty"`
	Gender        string `json:"gender,omitempty"`
	Income        string `json:"income,omitempty"`
	MaritalStatus string `json:"maritalStatus,omitempty"`
}

type Campaign struct {
	Account     *CampaignAccount  `json:"account"`
	Budget      *CampaignBudget   `json:"budget"`
	CreatedAt   *time.Time        `json:"createdAt,omitempty"`
	Description string            `json:"description,omitempty"`
	ExternalIDs []string          `json:"externalIDs,omitempty"`
	Flights     []CampaignFlight  `json:"flights"`
	ID          string            `json:"campaignID"`
	Name        string            `json:"name"`
	Objectives  []string          `json:"objectives,omitempty"`
	Status      *CampaignStatus   `json:"status,omitempty"`
	Taxonomy    *CampaignTaxonomy `json:"taxonomy,omitempty"`
	UpdatedAt   *time.Time        `json:"updatedAt,omitempty"`
}

type CampaignAccount struct {
	AccountExecutive *CampaignEntity          `json:"accountExecutive,omitempty"`
	Advertiser       *CampaignEntity          `json:"advertiser,omitempty"`
	Contacts         []CampaignAccountContact `json:"contacts,omitempty"`
	BillTo           *CampaignEntity          `json:"billTo,omitempty"`
	Buyer            *CampaignEntity          `json:"buyer,omitempty"`
	ThirdParty       *CampaignEntity          `json:"thirdParty,omitempty"`
}

type CampaignAccountContact struct {
	Email       string   `json:"email,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Primary     bool     `json:"primary,omitempty"`
	Title       string   `json:"title,omitempty"`
}

type CampaignBudget struct {
	Currency string  `json:"currency,omitempty"`
	Total    float32 `json:"total,omitempty"`
}

type CampaignEntity struct {
	CustomerID     string   `json:"customerID,omitempty"`
	CustomerNumber string   `json:"customerNumber,omitempty"`
	ExternalIDs    []string `json:"externalIDs,omitempty"`
	Name           string   `json:"name,omitempty"`
	Type           string   `json:"type,omitempty"`
}

type CampaignFlight struct {
	EndAt         *time.Time `json:"endAt,omitempty"`
	MediaProducts []struct {
		ExternalIDs []string `json:"externalIDs,omitempty"`
		ID          string   `json:"productID,omitempty"`
	} `json:"mediaProducts"`
	StartAt *time.Time `json:"startAt,omitempty"`
	Target  struct {
		Audience Audience `json:"audience,omitempty"`
		Location Location `json:"location,omitempty"`
	} `json:"target,omitempty"`
}

type CampaignStatus struct {
	CancelledAt          *time.Time `json:"cancelledAt,omitempty"`
	CancellationReason   string     `json:"cancellationReason,omitempty"`
	ContractCloseAt      *time.Time `json:"contractCloseAt,omitempty"`
	CreateEventEmittedAt *time.Time `json:"createEventEmittedAt,omitempty"`
	HardTakedownAt       *time.Time `json:"hardTakedownAt,omitempty"`
	ResponseDueBy        *time.Time `json:"responseDueBy,omitempty"`
}

type CampaignTaxonomy struct {
	CCO *CampaignTaxonomyCCO `json:"cco,omitempty"`
	IAB *CampaignTaxonomyIAB `json:"iab,omitempty"`
}

type CampaignTaxonomyCCO struct {
	Code     string `json:"code,omitempty"`
	FullCode string `json:"fullCode,omitempty"`
}

type CampaignTaxonomyIAB struct {
	V1 string `json:"v1,omitempty"`
	V2 string `json:"v2,omitempty"`
	V3 string `json:"v3,omitempty"`
}

type Location struct {
	Markets []struct {
		ExternalIDs []string `json:"externalIDs,omitempty"`
		ID          string   `json:"marketID,omitempty"`
	} `json:"markets,omitempty"`
	PointsOfInterest []struct {
		Coordinates []float32 `json:"coords,omitempty"`
		Radius      float32   `json:"radius,omitempty"`
	} `json:"pointsOfInterest,omitempty"`
	ZipCodes []string `json:"zipCodes,omitempty"`
}

/* Plan */

type Plan struct {
	CampaignID  string    `json:"campaignID"`
	CreatedAt   time.Time `json:"createdAt"`
	EndDate     time.Time `json:"endDate,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty"`
	ID          string    `json:"planID"`
	Markets     []Market  `json:"markets,omitempty"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"startDate,omitempty"`
	Status      Status    `json:"status,omitempty"`
	Type        string    `json:"type,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Market struct {
	Code        string   `json:"code,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty"`
}

type Status struct {
	Canceled *bool `json:"canceled,omitempty"`
}

/* PlanItem */

type PlanItem struct {
	BuyType        *BuyType        `json:"buyType,omitempty"`
	CreatedAt      time.Time       `json:"createdAt"`
	DigitalDetails *DigitalDetails `json:"digitalDetails,omitempty"`
	EndDate        time.Time       `json:"endDate,omitempty"`
	ExternalIDs    []string        `json:"externalIDs,omitempty"`
	ID             string          `json:"planItemID"`
	IsDigital      bool            `json:"isDigital,omitempty"`
	PlanID         string          `json:"planID"`
	PrintDetails   *PrintDetails   `json:"printDetails,omitempty"`
	StartDate      time.Time       `json:"startDate,omitempty"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type BuyType struct {
	Deliverable Deliverable `json:"deliverable,omitempty"`
	Flexibility Flexibility `json:"flexibility,omitempty"`
}

type DigitalDetails struct {
	ExternalIDs          []string `json:"externalIDs,omitempty"`
	NumberOfSpotsPerLoop int      `json:"numberOfSpotsPerLoop,omitempty"`
	Quantity             int      `json:"quantity,omitempty"`
	SpotLength           int      `json:"spotLength,omitempty"`
}

type PrintDetails struct {
	DisplayID   string   `json:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
}
