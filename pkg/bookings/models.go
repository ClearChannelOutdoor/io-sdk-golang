package bookings

import (
	"time"
)

type Booking struct {
	ID           string           `json:"bookingID"`
	BuyType      *BuyType         `json:"buyType,omitempty"`
	Cost         *float32         `json:"cost,omitempty"`
	CreatedAt    *time.Time       `json:"createdAt"`
	DeletedAt    *time.Time       `json:"deletedAt,omitempty"`
	Display      *DisplayDetails  `json:"display,omitempty"`
	EndDate      *time.Time       `json:"endDate,omitempty"`
	ExternalIDs  []string         `json:"externalIDs,omitempty"`
	Filler       bool             `json:"filler,omitempty"`
	Market       *Market          `json:"market,omitempty"`
	MediaProduct *MediaProduct    `json:"mediaProduct,omitempty"`
	Network      *NetworkDetails  `json:"network,omitempty"`
	OrderID      string           `json:"orderID,omitempty"`
	Quantity     *QuantityDetails `json:"quantity,omitempty"`
	StartDate    *time.Time       `json:"startDate,omitempty"`
	Status       *BookingStatus   `json:"status,omitempty"`
	UpdatedAt    *time.Time       `json:"updatedAt"`
	Waitlisted   *bool            `json:"waitlisted,omitempty"`
}
type BookingStatus struct {
	Status BookingStatusValue   `json:"status,omitempty"`
	Sync   map[string]time.Time `json:"sync,omitempty"`
}

type BuyType struct {
	Deliverable      Deliverable      `json:"deliverable,omitempty"`
	Flexibility      Flexibility      `json:"flexibility,omitempty"`
	RevenueSpecifier RevenueSpecifier `json:"revenueSpecifier,omitempty"`
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

type DisplayDetails struct {
	DisplayID   string   `json:"displayID,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	Slots       int      `json:"slots,omitempty"`
}

type Market struct {
	Code        string   `json:"code,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
	MarketID    string   `json:"marketID,omitempty"`
}

type MediaProduct struct {
	ProductCode string `json:"productCode,omitempty"`
	TypeCode    string `json:"typeCode,omitempty"`
}

type NetworkDetails struct {
	DailyEndTime      string             `json:"dailyEndTime,omitempty"`
	DailyStartTime    string             `json:"dailyStartTime,omitempty"`
	DaysToPlay        *DigitalDaysToPlay `json:"daysToPlay,omitempty"`
	Frequency         int                `json:"frequency,omitempty"`
	NetworkID         string             `json:"networkID,omitempty"`
	NumberOfSlots     int                `json:"numberOfSlots,omitempty"`
	SlotSeconds       float32            `json:"slotSeconds,omitempty"`
	SlotSlices        int                `json:"slotSlices,omitempty"`
	SpecificStartTime string             `json:"specificStartTime,omitempty"`
}

type QuantityCustomDetails struct {
	ExternalIDs []string          `json:"externalIDs,omitempty"`
	Displays    []*DisplayDetails `json:"displays,omitempty"`
}

type QuantityDetails struct {
	Custom            *QuantityCustomDetails     `json:"custom,omitempty"`
	FullMarket        *QuantityFullMarketDetails `json:"fullMarket,omitempty"`
	RequestedQuantity int                        `json:"requestedQuantity,omitempty"`
}

type QuantityFullMarketDetails struct {
	Weight int `json:"weight,omitempty"`
}

const ExplicitEmpty = ""

type BookingStatusValue string

const (
	Draft    BookingStatusValue = "Draft"
	Reserved BookingStatusValue = "Reserved"
	Booked   BookingStatusValue = "Booked"
	Canceled BookingStatusValue = "Canceled"
)

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
	EmptyFlexibility Flexibility = ExplicitEmpty
)

type RevenueSpecifier string

const (
	Bonus                 RevenueSpecifier = "Bonus"
	EmptyRevenueSpecifier RevenueSpecifier = ExplicitEmpty
	FBI                   RevenueSpecifier = "FBI"
	Lessor                RevenueSpecifier = "Lessor"
	MakeGood              RevenueSpecifier = "Make Good"
	PSA                   RevenueSpecifier = "PSA"
	RFR                   RevenueSpecifier = "Right of First Refusal"
	Trade                 RevenueSpecifier = "Trade"
)
