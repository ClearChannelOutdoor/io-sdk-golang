package bookings

import "time"

type Booking struct {
	Canceled    *bool           `json:"canceled,omitempty" bson:"canceled,omitempty"`
	CreatedAt   time.Time       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Digital     *DigitalDetails `json:"digital,omitempty" bson:"digital,omitempty"`
	EndDate     *time.Time      `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExternalIDs []string        `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Filler      bool            `json:"filler,omitempty" bson:"filler,omitempty"`
	ID          string          `json:"bookingID,omitempty" bson:"bookingID,omitempty"`
	OrderLineID string          `json:"orderLineID,omitempty" bson:"orderLineID,omitempty"`
	Print       *PrintDetails   `json:"print,omitempty" bson:"print,omitempty"`
	StartDate   *time.Time      `json:"startDate,omitempty" bson:"startDate,omitempty"`
	UpdatedAt   time.Time       `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Waitlisted  *bool           `json:"waitlisted,omitempty" bson:"waitlisted,omitempty"`
}

type PrintDetails struct {
	DisplayID   string   `json:"displayID,omitempty" bson:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type DigitalDetails struct {
	DailyEndTime      string             `json:"dailyEndTime,omitempty" bson:"dailyEndTime,omitempty"`
	DailyStartTime    string             `json:"dailyStartTime,omitempty" bson:"dailyStartTime,omitempty"`
	DaysToPlay        *DigitalDaysToPlay `json:"daysToPlay,omitempty" bson:"daysToPlay,omitempty"`
	Frequency         int                `json:"frequency,omitempty" bson:"frequency,omitempty"`
	NetworkID         string             `json:"networkID,omitempty" bson:"networkID,omitempty"`
	NumberOfSlots     int                `json:"numberOfSlots,omitempty" bson:"numberOfSlots,omitempty"`
	SlotSeconds       float32            `json:"slotSeconds,omitempty" bson:"slotSeconds,omitempty"`
	SlotSlices        int                `json:"slotSlices,omitempty" bson:"slotSlices,omitempty"`
	SpecificStartTime string             `json:"specificStartTime,omitempty" bson:"specificStartTime,omitempty"`
}

type DigitalDaysToPlay struct {
	Sunday    bool `json:"sunday" bson:"sunday"`
	Monday    bool `json:"monday" bson:"monday"`
	Tuesday   bool `json:"tuesday" bson:"tuesday"`
	Wednesday bool `json:"wednesday" bson:"wednesday"`
	Thursday  bool `json:"thursday" bson:"thursday"`
	Friday    bool `json:"friday" bson:"friday"`
	Saturday  bool `json:"saturday" bson:"saturday"`
}
