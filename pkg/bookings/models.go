package bookings

import (
	"time"
)

type (
	BookingStatus  string
	LifecycleEvent string
)

const (
	Draft    BookingStatus = "Draft"
	Reserved BookingStatus = "Reserved"
	Booked   BookingStatus = "Booked"
	Canceled BookingStatus = "Canceled"
)

type Booking struct {
	ID      string   `json:"bookingID,omitempty" bson:"bookingID,omitempty"`
	BuyType *BuyType `json:"buyType,omitempty" bson:"buyType,omitempty"`
	// todo: remove buyTypeID once dependency on buy-type-api is removed from all services
	BuyTypeID    string           `json:"buyTypeID,omitempty" bson:"buyTypeID,omitempty"`
	Cost         *float32         `json:"cost,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	DeletedAt    *time.Time       `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
	Digital      *DigitalDetails  `json:"digital,omitempty" bson:"digital,omitempty"`
	EndDate      time.Time        `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExternalIDs  []string         `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Filler       bool             `json:"filler,omitempty" bson:"filler,omitempty"`
	Lifecycle    BookingLifecycle `json:"lifecycle,omitempty" bson:"lifecycle,omitempty"`
	MarketID     string           `json:"marketID,omitempty" bson:"marketID,omitempty"`
	MediaProduct MediaProduct     `json:"mediaProduct,omitempty" bson:"mediaProduct,omitempty"`
	OrderID      string           `json:"orderID,omitempty" bson:"orderID,omitempty"`
	Print        *PrintDetails    `json:"print,omitempty" bson:"print,omitempty"`
	StartDate    time.Time        `json:"startDate,omitempty" bson:"startDate,omitempty"`
	Status       BookingStatus    `json:"status,omitempty" bson:"status,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Waitlisted   *bool            `json:"waitlisted,omitempty" bson:"waitlisted,omitempty"`
}

type BuyType struct {
	Deliverable      Deliverable      `json:"deliverable,omitempty" bson:"deliverable,omitempty"`
	Flexibility      Flexibility      `json:"flexibility,omitempty" bson:"flexibility,omitempty"`
	RevenueSpecifier RevenueSpecifier `json:"revenueSpecifier,omitempty" bson:"revenueSpecifier,omitempty"`
}

const ExplicitEmpty = ""

type Deliverable string

const (
	// unique values
	Display     Deliverable = "Display"
	Impressions Deliverable = "Impressions"
	Override    Deliverable = "Override"
	Quantity    Deliverable = "Quantity"
)

type Flexibility string

const (
	Fixed            Flexibility = "Fixed"
	Flexible         Flexibility = "Flexible"
	EmptyFlexibility Flexibility = ExplicitEmpty // for bonus override only
)

type RevenueSpecifier string

const (
	Bonus                 RevenueSpecifier = "Bonus"
	EmptyRevenueSpecifier RevenueSpecifier = ExplicitEmpty
	FBI                   RevenueSpecifier = "FBI"
	Lessor                RevenueSpecifier = "Lessor"
	MakeGood              RevenueSpecifier = "Make Good"
	PSA                   RevenueSpecifier = "PSA" // lives on order, not booking, may still come through on an orderline message and will need to be parsed correctly
	RFR                   RevenueSpecifier = "Right of First Refusal"
	Trade                 RevenueSpecifier = "Trade"
)

type MediaProduct struct {
	ProductCode string `json:"productCode,omitempty" bson:"productCode,omitempty"`
	TypeCode    string `json:"typeCode,omitempty" bson:"typeCode,omitempty"`
}

type BookingLifecycle struct {
	QuattroCreate       LifecycleTimestamp `json:"quattroCreate,omitempty" bson:"quattroCreate,omitempty"`
	QuattroReserve      LifecycleTimestamp `json:"quattroReserve,omitempty" bson:"quattroReserve,omitempty"`
	QuattroBook         LifecycleTimestamp `json:"quattroBook,omitempty" bson:"quattroBook,omitempty"`
	QuattroDelete       LifecycleTimestamp `json:"quattroDelete,omitempty" bson:"quattroDelete,omitempty"`
	QuattroResetToDraft LifecycleTimestamp `json:"quattroResetToDraft,omitempty" bson:"quattroResetToDraft,omitempty"`
}

type LifecycleTimestamp struct {
	Sent     time.Time `json:"sent,omitempty" bson:"sent,omitempty"`
	Received time.Time `json:"received,omitempty" bson:"received,omitempty"`
}

type SegmentDetails struct {
	DetailCode  string `json:"detailCode,omitempty" bson:"detailCode,omitempty"`
	SegmentCode string `json:"segmentCode,omitempty" bson:"segmentCode,omitempty"`
	TRP         *int   `json:"trp,omitempty" bson:"trp,omitempty"`
}

type PrintDetails struct {
	DisplayID   string   `json:"displayID,omitempty" bson:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
}

type DigitalDetails struct {
	NetworkID         string             `json:"networkID,omitempty" bson:"networkID,omitempty"`
	DailyStartTime    string             `json:"dailyStartTime,omitempty" bson:"dailyStartTime,omitempty"`
	DailyEndTime      string             `json:"dailyEndTime,omitempty" bson:"dailyEndTime,omitempty"`
	DaysToPlay        *DigitalDaysToPlay `json:"daysToPlay,omitempty" bson:"daysToPlay,omitempty"`
	NumberOfSlots     int                `json:"numberOfSlots,omitempty" bson:"numberOfSlots,omitempty"`
	Frequency         int                `json:"frequency,omitempty" bson:"frequency,omitempty"`
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
