package measures

import "time"

type Construction struct {
	Classification string `json:"classification,omitempty"`
	Placement      string `json:"placement,omitempty"`
	Type           string `json:"type,omitempty"`
}

type Frame struct {
	Geopath struct {
		FrameID int        `json:"frameID,omitempty"`
		Created *time.Time `json:"createdAt,omitempty"`
		Updated *time.Time `json:"updatedAt,omitempty"`
	} `json:"geopath,omitempty"`
	Construction    *Construction `json:"construction,omitempty"`
	Illumination    *Illumination `json:"illumination,omitempty"`
	IsAudio         bool          `json:"isAudio,omitempty"`
	IsDigital       bool          `json:"isDigital,omitempty"`
	IsFullMotion    bool          `json:"isFullMotion,omitempty"`
	IsInteractive   bool          `json:"isInteractive,omitempty"`
	IsPartialMotion bool          `json:"isPartialMotion,omitempty"`
	IsRotating      bool          `json:"isRotating,omitempty"`
	Location        *Location     `json:"location,omitempty"`
	Media           *Media        `json:"media,omitempty"`
	Operator        *Operator     `json:"operator,omitempty"`
}

type FramesStatus struct {
	UpdatedAt *time.Time `json:"updated_at"`
}

type Illumination struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Type      string `json:"type,omitempty"`
}

type Location struct {
	CBSA          *CBSA      `json:"cbsa,omitempty"`
	Coords        *[]float32 `json:"coords,omitempty"`
	County        string     `json:"county,omitempty"`
	DMA           *DMA       `json:"dma,omitempty"`
	Level         int        `json:"level,omitempty"`
	LevelsVisible string     `json:"levelsVisible,omitempty"`
	Orientation   float32    `json:"orientation,omitempty"`
	Places        *[]Place   `json:"places,omitempty"`
	PrimaryArtery string     `json:"primaryArtery,omitempty"`
	PrimaryRead   string     `json:"primaryRead,omitempty"`
	State         string     `json:"state,omitempty"`
	Timezone      string     `json:"timezone,omitempty"`
	Type          string     `json:"type,omitempty"`
	ZipCode       string     `json:"zipCode,omitempty"`
}

type Measure struct {
	Geopath struct {
		Created            time.Time `json:"createdAt"`
		FrameID            int       `json:"frameID"`
		FrameLength        int       `json:"frameLength"`
		FrameShareOfVoice  float32   `json:"frameShareOfVoice"`
		LayoutShareOfVoice float32   `json:"layoutShareOfVoice"`
		ProductName        string    `json:"productName"`
		Updated            time.Time `json:"updatedAt"`
	} `json:"geopath"`
	Calculated time.Time `json:"calculatedAt"` // time calculated date/time of the spot
	Frequency  struct {
		Average      float32 `json:"average"`      // calculated average frequency of the target audience
		EffectiveAvg float32 `json:"effectiveAvg"` // calculated average frequency of persons reached at least effectiveMin times
		EffectiveMin int     `json:"effectiveMin"` // min frequency to be considered in reach; defaults to 1
	} `json:"frequency"`
	Hourly *struct {
		Monday    schema.HourlyImpression `json:"mon,omitempty"`
		Tuesday   schema.HourlyImpression `json:"tue,omitempty"`
		Wednesday schema.HourlyImpression `json:"wed,omitempty"`
		Thursday  schema.HourlyImpression `json:"thu,omitempty"`
		Friday    schema.HourlyImpression `json:"fri,omitempty"`
		Saturday  schema.HourlyImpression `json:"sat,omitempty"`
		Sunday    schema.HourlyImpression `json:"sun,omitempty"`
	} `json:"hourly,omitempty"`
	Impressions struct {
		Market  float32 `json:"market"`  // impressions reaching the geographic market (i.e. Global, DMA or CBSA)
		Segment float32 `json:"segment"` // impressesions reaching target audience (segment) overall
		Total   float32 `json:"total"`   // total impressions
	} `json:"impressions"`
	Population struct {
		Market  int `json:"market"`  // total population in the geographic market (i.e. Global, DMA or CBSA)
		Segment int `json:"segment"` // target segment population
	} `json:"population"`
	Reach struct {
		Effective int `json:"effective"` // total target in-market population reached at least eff_freq_min times
		Total     int `json:"total"`     // number of people in the geographic market that are reached
	} `json:"reach"`
	Segment struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"segment"`
}
