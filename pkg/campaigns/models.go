package campaigns

import "time"

type audience struct {
	Demographics struct {
		Age           string `json:"age,omitempty"`
		Ethnicity     string `json:"ethnicity,omitempty"`
		Gender        string `json:"gender,omitempty"`
		Income        string `json:"income,omitempty"`
		MaritalStatus string `json:"maritalStatus,omitempty"`
	} `json:"demographics,omitempty"`
}

type Campaign struct {
	Account         CampaignAccount   `json:"account"`
	Budget          CampaignBudget    `json:"budget"`
	ContractCloseAt time.Time         `json:"contractCloseAt,omitempty"`
	CreatedAt       time.Time         `json:"createdAt,omitempty"`
	Description     string            `json:"description,omitempty"`
	ExternalIDs     []string          `json:"externalIDs,omitempty"`
	Flight          CampaignFlight    `json:"flight"`
	Goals           []string          `json:"goals,omitempty"`
	ID              string            `json:"campaignID"`
	Name            string            `json:"name"`
	Segments        []CampaignSegment `json:"segments"`
	Taxonomy        struct {
		CCO struct {
			Code     string `json:"code,omitempty"`
			FullCode string `json:"fullCode,omitempty"`
		} `json:"cco,omitempty"`
		IAB struct {
			V1 string `json:"v1,omitempty"`
			V2 string `json:"v2,omitempty"`
		} `json:"iab,omitempty"`
	} `json:"taxonomy,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CampaignAccount struct {
	Advertiser customerEntity `json:"advertiser,omitempty"`
	BillTo     customerEntity `json:"billTo,omitempty"`
	Buyer      customerEntity `json:"buyer,omitempty"`
	ThirdParty customerEntity `json:"thirdParty,omitempty"`
}

type CampaignBudget struct {
	Currency string  `json:"currency,omitempty"`
	Total    float32 `json:"total,omitempty"`
}

type CampaignFlight struct {
	EndAt   *time.Time `json:"endAt,omitempty"`
	StartAt *time.Time `json:"startAt,omitempty"`
}

type CampaignSegment struct {
	Flight   CampaignFlight `json:"flight,omitempty"`
	Products []struct {
		ExternalIDs []string `json:"externalIDs,omitempty"`
		ID          string   `json:"productID,omitempty"`
	} `json:"products"`
	Target struct {
		Audience audience `json:"audience,omitempty"`
		Location location `json:"location,omitempty"`
	} `json:"target,omitempty"`
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
