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
	ID      string   `json:"bookingID,omitempty"`
	BuyType *BuyType `json:"buyType,omitempty"`
	// todo: remove buyTypeID once dependency on buy-type-api is removed from all services
	BuyTypeID    string           `json:"buyTypeID,omitempty"`
	Cost         *float32         `json:"cost,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	DeletedAt    *time.Time       `json:"deletedAt,omitempty"`
	Digital      *DigitalDetails  `json:"digital,omitempty"`
	EndDate      time.Time        `json:"endDate,omitempty"`
	ExternalIDs  []string         `json:"externalIDs,omitempty"`
	Filler       bool             `json:"filler,omitempty"`
	Lifecycle    BookingLifecycle `json:"lifecycle,omitempty"`
	MarketID     string           `json:"marketID,omitempty"`
	MediaProduct MediaProduct     `json:"mediaProduct,omitempty"`
	OrderID      string           `json:"orderID,omitempty"`
	Print        *PrintDetails    `json:"print,omitempty"`
	StartDate    time.Time        `json:"startDate,omitempty"`
	Status       BookingStatus    `json:"status,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty"`
	Waitlisted   *bool            `json:"waitlisted,omitempty"`
}

type BuyType struct {
	Deliverable      Deliverable      `json:"deliverable,omitempty"`
	Flexibility      Flexibility      `json:"flexibility,omitempty"`
	RevenueSpecifier RevenueSpecifier `json:"revenueSpecifier,omitempty"`
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
	ProductCode string `json:"productCode,omitempty"`
	TypeCode    string `json:"typeCode,omitempty"`
}

type BookingLifecycle struct {
	QuattroCreate       LifecycleTimestamp `json:"quattroCreate,omitempty"`
	QuattroReserve      LifecycleTimestamp `json:"quattroReserve,omitempty"`
	QuattroBook         LifecycleTimestamp `json:"quattroBook,omitempty"`
	QuattroDelete       LifecycleTimestamp `json:"quattroDelete,omitempty"`
	QuattroResetToDraft LifecycleTimestamp `json:"quattroResetToDraft,omitempty"`
}

type LifecycleTimestamp struct {
	Sent     time.Time `json:"sent,omitempty"`
	Received time.Time `json:"received,omitempty"`
}

type SegmentDetails struct {
	DetailCode  string `json:"detailCode,omitempty"`
	SegmentCode string `json:"segmentCode,omitempty"`
	TRP         *int   `json:"trp,omitempty"`
}

type PrintDetails struct {
	DisplayID   string   `json:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
}

type DigitalDetails struct {
	NetworkID         string             `json:"networkID,omitempty"`
	DailyStartTime    string             `json:"dailyStartTime,omitempty"`
	DailyEndTime      string             `json:"dailyEndTime,omitempty"`
	DaysToPlay        *DigitalDaysToPlay `json:"daysToPlay,omitempty"`
	NumberOfSlots     int                `json:"numberOfSlots,omitempty"`
	Frequency         int                `json:"frequency,omitempty"`
	SlotSeconds       float32            `json:"slotSeconds,omitempty"`
	SlotSlices        int                `json:"slotSlices,omitempty"`
	SpecificStartTime string             `json:"specificStartTime,omitempty"`
}

type DigitalDaysToPlay struct {
	Sunday    bool `json:"sunday"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
}
