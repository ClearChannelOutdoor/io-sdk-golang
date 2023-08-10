package campaigns

import "time"

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
	Advertiser customerEntity           `json:"advertiser,omitempty"`
	Contacts   []CampaignAccountContact `json:"contacts,omitempty"`
	BillTo     customerEntity           `json:"billTo,omitempty"`
	Buyer      customerEntity           `json:"buyer,omitempty"`
	ThirdParty customerEntity           `json:"thirdParty,omitempty"`
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
	CancelledAt     *time.Time `json:"cancelledAt,omitempty" bson:"cancelledAt,omitempty"`
	ContractCloseAt *time.Time `json:"contractCloseAt,omitempty" bson:"contractCloseAt,omitempty"`
	HardTakedownAt  *time.Time `json:"hardTakedownAt,omitempty" bson:"hardTakedownAt,omitempty"`
	ResponseBy      *time.Time `json:"responseBy,omitempty" bson:"responseBy,omitempty"`
}

type customerEntity struct {
	CustomerID  string   `json:"customerID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Name        string   `json:"name,omitempty"`
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

type Plan struct {
	CampaignID string          `json:"campaignID"`
	CreatedAt  time.Time       `json:"createdAt,omitempty"`
	ID         string          `json:"planID"`
	Inventory  []PlanInventory `json:"inventory"`
	UpdatedAt  time.Time       `json:"updatedAt,omitempty"`
}

type PlanInventory struct {
	DisplayID string         `json:"displayID"`
	Flight    CampaignFlight `json:"flight,omitempty"`
	ProductID string         `json:"productID"`
}
