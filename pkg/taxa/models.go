package taxa

type CCOCode struct {
	Code     string       `json:"code"`
	FullCode string       `json:"fullCode"`
	Name     string       `json:"name"`
	Unmapped bool         `json:"unmapped"`
	V1       *V1Taxonomy  `json:"v1,omitempty"`
	V2       *IABTaxonomy `json:"v2,omitempty"`
	V3       *IABTaxonomy `json:"v3,omitempty"`
}

type V1Taxonomy struct {
	IABCode    string `json:"iabCode"`
	Category   string `json:"category"`
	Tier1      string `json:"tier1,omitempty"`
	Deprecated bool   `json:"deprecated"`
}

type IABTaxonomy struct {
	UniqueID  string `json:"uniqueID"`
	Parent    string `json:"parent,omitempty"`
	Name      string `json:"name"`
	Tier1     string `json:"tier1"`
	Tier2     string `json:"tier2,omitempty"`
	Tier3     string `json:"tier3,omitempty"`
	Tier4     string `json:"tier4,omitempty"`
	Extension string `json:"extension,omitempty"`
}
