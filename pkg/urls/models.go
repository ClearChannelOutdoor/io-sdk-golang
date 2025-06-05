package urls

import "time"

type Location struct {
	CreatedAt *time.Time      `json:"createdAt,omitempty"`
	Slugs     []*LocationSlug `json:"slugs,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	URL       *string         `json:"url,omitempty"`
}

type LocationSlug struct {
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	LastRequestedAt *time.Time `json:"lastRequestedAt,omitempty"`
	ShortURL        string     `json:"shortURL,omitempty"`
	Slug            *string    `json:"slug,omitempty"`
	Visits          *uint      `json:"visits,omitempty"`
}
