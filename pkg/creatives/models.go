package creatives

import "time"

type AdCreative struct {
	Asset       AdCreativeAsset `json:"asset"`
	CreatedAt   time.Time       `json:"createdAt"`
	Description string          `json:"description,omitempty"`
	ExternalIDs []string        `json:"externalIDs,omitempty"`
	FileType    string          `json:"fileType,omitempty"`
	Format      string          `json:"format,omitempty"`
	Hidden      *bool           `json:"hidden,omitempty"`
	ID          string          `json:"creativeID"`
	SubmittedBy string          `json:"submittedBy,omitempty"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

type AdCreativeAsset struct {
	PreviewURL   string `json:"previewURL,omitempty"`
	ThumbnailURL string `json:"thumbnailURL,omitempty"`
	URL          string `json:"url,omitempty"`
}

type Creative struct {
	Asset       CreativeAsset `json:"asset"`
	CreatedAt   time.Time     `json:"createdAt"`
	Description string        `json:"description,omitempty"`
	ExternalIDs []string      `json:"externalIDs,omitempty"`
	FileType    string        `json:"fileType,omitempty"`
	Format      string        `json:"format,omitempty"`
	Hidden      bool          `json:"hidden,omitempty"`
	ID          string        `json:"creativeID"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type CreativeAsset struct {
	PreviewURL   string `json:"previewURL,omitempty"`
	ThumbnailURL string `json:"thumbnailURL,omitempty"`
	URL          string `json:"url,omitempty"`
}
