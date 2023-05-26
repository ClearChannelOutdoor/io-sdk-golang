package taxa

type CCOCode struct {
	Code     string         `json:"code"`
	FullCode string         `json:"fullCode"`
	Name     string         `json:"name"`
	Unmapped bool           `json:"unmapped"`
	V1       *IABV1Taxonomy `json:"v1,omitempty"`
}

type IABV1Taxonomy struct {
	RTBID      string         `json:"rtbID"`
	Subject    string         `json:"subject"`
	Deprecated bool           `json:"deprecated"`
	V2         *IABV2Taxonomy `json:"v2,omitempty"`
}

type IABV2Taxonomy struct {
	UniqueID  string `json:"uniqueID"`
	Parent    string `json:"parent,omitempty"`
	Name      string `json:"name"`
	Tier1     string `json:"tier1"`
	Tier2     string `json:"tier2,omitempty"`
	Tier3     string `json:"tier3,omitempty"`
	Tier4     string `json:"tier4,omitempty"`
	Extension string `json:"extension,omitempty"`
}
