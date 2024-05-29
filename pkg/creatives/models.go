package creatives

import "time"

type AdCreative struct {
	ID          string          `json:"creativeID" bson:"_id"`
	Asset       AdCreativeAsset `json:"asset" bson:"asset"`
	CreatedAt   time.Time       `json:"createdAt" bson:"createdAt"`
	Description string          `json:"description,omitempty" bson:"description,omitempty"`
	ExternalIDs []string        `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	FileType    string          `json:"fileType,omitempty" bson:"fileType,omitempty"`
	Format      string          `json:"format,omitempty" bson:"format,omitempty"`
	Hidden      *bool           `json:"hidden,omitempty" bson:"hidden,omitempty"`
	SubmittedBy string          `json:"submittedBy,omitempty" bson:"submittedBy,omitempty"`
	UpdatedAt   time.Time       `json:"updatedAt" bson:"updatedAt"`
}

type AdCreativeAsset struct {
	PreviewURL   string `json:"previewURL,omitempty" bson:"previewURL,omitempty"`
	ThumbnailURL string `json:"thumbnailURL,omitempty" bson:"thumbnailURL,omitempty"`
	URL          string `json:"url,omitempty" bson:"url,omitempty"`
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
