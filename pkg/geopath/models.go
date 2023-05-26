package geopath

import "time"

type CBSA struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type DMA struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Construction struct {
	Classification string `json:"classification,omitempty"`
	Placement      string `json:"placement,omitempty"`
	Type           string `json:"type,omitempty"`
}

type ConstructionClassification struct {
	Classification string `json:"construction.classification"`
	FrameCount     int    `json:"count"`
}

type ConstructionPlacement struct {
	Placement  string `json:"construction.placement"`
	FrameCount int    `json:"count"`
}

type ConstructionType struct {
	Type       string `json:"construction.type"`
	FrameCount int    `json:"count"`
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

type HourlyImpression struct {
	Hour0  float32 `json:"0"`
	Hour1  float32 `json:"1"`
	Hour2  float32 `json:"2"`
	Hour3  float32 `json:"3"`
	Hour4  float32 `json:"4"`
	Hour5  float32 `json:"5"`
	Hour6  float32 `json:"6"`
	Hour7  float32 `json:"7"`
	Hour8  float32 `json:"8"`
	Hour9  float32 `json:"9"`
	Hour10 float32 `json:"10"`
	Hour11 float32 `json:"11"`
	Hour12 float32 `json:"12"`
	Hour13 float32 `json:"13"`
	Hour14 float32 `json:"14"`
	Hour15 float32 `json:"15"`
	Hour16 float32 `json:"16"`
	Hour17 float32 `json:"17"`
	Hour18 float32 `json:"18"`
	Hour19 float32 `json:"19"`
	Hour20 float32 `json:"20"`
	Hour21 float32 `json:"21"`
	Hour22 float32 `json:"22"`
	Hour23 float32 `json:"23"`
}

type Illumination struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Type      string `json:"type,omitempty"`
}

type IlluminationType struct {
	Type       string `json:"illumination.type"`
	FrameCount int    `json:"count"`
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

type LocationType struct {
	Type       string `json:"location.type"`
	FrameCount int    `json:"count"`
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
		Monday    HourlyImpression `json:"mon,omitempty"`
		Tuesday   HourlyImpression `json:"tue,omitempty"`
		Wednesday HourlyImpression `json:"wed,omitempty"`
		Thursday  HourlyImpression `json:"thu,omitempty"`
		Friday    HourlyImpression `json:"fri,omitempty"`
		Saturday  HourlyImpression `json:"sat,omitempty"`
		Sunday    HourlyImpression `json:"sun,omitempty"`
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

type MeasuresStatus struct {
	UpdatedAt *time.Time `json:"updated_at"`
}

type Media struct {
	MaxHeight int    "json:\"maxHeight,omitempty\""
	MaxWidth  int    "json:\"maxWidth,omitempty\""
	Name      string "json:\"name,omitempty\""
	Status    string "json:\"status,omitempty\""
	Type      string "json:\"type,omitempty\""
}

type MediaType struct {
	Type       string `json:"media.type"`
	FrameCount int    `json:"count"`
}

type Operator struct {
	Name           string `json:"name,omitempty"`
	Code           string `json:"code,omitempty"`
	ParentName     string `json:"parentName,omitempty"`
	Representation string `json:"representation,omitempty"`
}

type Place struct {
	City          string `json:"city,omitempty"`
	Name          string `json:"name,omitempty"`
	ShortName     string `json:"shortName,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
	Type          string `json:"type,omitempty"`
	UpdatedDesc   string `json:"updatedDesc,omitempty"`
}

type SegmentID struct {
	ID    int `json:"segment.id"`
	Count int `json:"count"`
}

type SegmentName struct {
	Name  string `json:"segment.name"`
	Count int    `json:"count"`
}
