package campaigns

import "time"

/* Campaign */
type audience struct {
	Demographics []struct {
		Age           string `json:"age,omitempty"`
		Ethnicity     string `json:"ethnicity,omitempty"`
		Gender        string `json:"gender,omitempty"`
		Income        string `json:"income,omitempty"`
		MaritalStatus string `json:"maritalStatus,omitempty"`
	} `json:"demographics,omitempty"`
}

type Campaign struct {
	Account     *CampaignAccount `json:"account"`
	Budget      *CampaignBudget  `json:"budget"`
	CreatedAt   *time.Time       `json:"createdAt,omitempty"`
	Description string           `json:"description,omitempty"`
	ExternalIDs []string         `json:"externalIDs,omitempty"`
	Flights     []CampaignFlight `json:"flights"`
	ID          string           `json:"campaignID"`
	Name        string           `json:"name"`
	Objectives  []string         `json:"objectives,omitempty"`
	Status      *CampaignStatus  `json:"status,omitempty" bson:"status,omitempty"`
	Taxonomy    *struct {
		CCO *struct {
			Code     string `json:"code,omitempty"`
			FullCode string `json:"fullCode,omitempty"`
		} `json:"cco,omitempty"`
		IAB *struct {
			V1 string `json:"v1,omitempty"`
			V2 string `json:"v2,omitempty"`
			V3 string `json:"v3,omitempty"`
		} `json:"iab,omitempty"`
	} `json:"taxonomy,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type CampaignAccount struct {
	AccountExecutive customerEntity           `json:"accountExecutive,omitempty"`
	Advertiser       customerEntity           `json:"advertiser,omitempty"`
	Contacts         []CampaignAccountContact `json:"contacts,omitempty"`
	BillTo           customerEntity           `json:"billTo,omitempty"`
	Buyer            customerEntity           `json:"buyer,omitempty"`
	ThirdParty       customerEntity           `json:"thirdParty,omitempty"`
}

type CampaignAccountContact struct {
	Email       string   `json:"email,omitempty" bson:"email,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
	Phone       string   `json:"phone,omitempty" bson:"phone,omitempty"`
	Primary     bool     `json:"primary,omitempty" bson:"primary,omitempty"`
	Title       string   `json:"title,omitempty" bson:"title,omitempty"`
}

type CampaignBudget struct {
	Currency string  `json:"currency,omitempty"`
	Total    float32 `json:"total,omitempty"`
}

type CampaignFlight struct {
	EndAt         *time.Time `json:"endAt,omitempty"`
	MediaProducts []struct {
		ExternalIDs []string `json:"externalIDs,omitempty"`
		ID          string   `json:"productID,omitempty"`
	} `json:"mediaProducts"`
	StartAt *time.Time `json:"startAt,omitempty"`
	Target  struct {
		Audience audience `json:"audience,omitempty"`
		Location location `json:"location,omitempty"`
	} `json:"target,omitempty"`
}

type CampaignStatus struct {
	CancelledAt        *time.Time `json:"cancelledAt,omitempty" bson:"cancelledAt,omitempty"`
	CancellationReason string     `json:"cancellationReason,omitempty" bson:"cancellationReason,omitempty"`
	ContractCloseAt    *time.Time `json:"contractCloseAt,omitempty" bson:"contractCloseAt,omitempty"`
	HardTakedownAt     *time.Time `json:"hardTakedownAt,omitempty" bson:"hardTakedownAt,omitempty"`
	ResponseBy         *time.Time `json:"responseBy,omitempty" bson:"responseBy,omitempty"`
}

type customerEntity struct {
	CustomerID     string   `json:"customerID,omitempty"`
	CustomerNumber string   `json:"customerNumber,omitempty"`
	ExternalIDs    []string `json:"externalIDs,omitempty"`
	Name           string   `json:"name,omitempty"`
	Type           string   `json:"type,omitempty"`
}

type location struct {
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
	CampaignID  string    `json:"campaignID" bson:"campaignID"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	EndDate     time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExternalIDs []string  `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID          string    `json:"planID" bson:"planID"`
	Markets     []Market  `json:"markets,omitempty" bson:"markets,omitempty"`
	Name        string    `json:"name" bson:"name"`
	StartDate   time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	Status      Status    `json:"status,omitempty" bson:"status,omitempty"`
	Type        string    `json:"type,omitempty" bson:"type,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Market struct {
	Code        string   `json:"code,omitempty" bson:"code,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
}

type Status struct {
	Canceled *bool `json:"canceled,omitempty" bson:"canceled,omitempty"`
}

/* PlanItem */

type PlanItem struct {
	BuyType        *BuyType        `json:"buyType,omitempty" bson:"buyType,omitempty"`
	CreatedAt      time.Time       `json:"createdAt" bson:"createdAt"`
	DigitalDetails *DigitalDetails `json:"digitalDetails,omitempty" bson:"digitalDetails,omitempty"`
	EndDate        time.Time       `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExternalIDs    []string        `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID             string          `json:"planItemID" bson:"planItemID"`
	IsDigital      bool            `json:"isDigital,omitempty" bson:"isDigital,omitempty"`
	PlanID         string          `json:"planID" bson:"planID"`
	PrintDetails   *PrintDetails   `json:"printDetails,omitempty" bson:"printDetails,omitempty"`
	StartDate      time.Time       `json:"startDate,omitempty" bson:"startDate,omitempty"`
	UpdatedAt      time.Time       `json:"updatedAt" bson:"updatedAt"`
}

type BuyType struct {
	Deliverable Deliverable `json:"deliverable,omitempty" bson:"deliverable,omitempty"`
	Flexibility Flexibility `json:"flexibility,omitempty" bson:"flexibility,omitempty"`
}

type DigitalDetails struct {
	ExternalIDs          []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	NumberOfSpotsPerLoop int      `json:"numberOfSpotsPerLoop,omitempty" bson:"numberOfSpotsPerLoop,omitempty"`
	Quantity             int      `json:"quantity,omitempty" bson:"quantity,omitempty"`
	SpotLength           int      `json:"spotLength,omitempty" bson:"spotLength,omitempty"`
}
type PrintDetails struct {
	DisplayID   string   `json:"displayID,omitempty" bson:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}
