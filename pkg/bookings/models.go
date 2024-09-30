package bookings

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type BookingStatus string
type LifecycleEvent string

const (
	Draft    BookingStatus = "Draft"
	Reserved BookingStatus = "Reserved"
	Booked   BookingStatus = "Booked"
	Canceled BookingStatus = "Canceled"

	CreateAsReserved LifecycleEvent = "CreateAsReserved"
	UpdateToBooked   LifecycleEvent = "UpdateToBooked"
)

type Booking struct {
	ID          string           `json:"bookingID,omitempty" bson:"bookingID,omitempty"`
	Canceled    *bool            `json:"canceled,omitempty" bson:"canceled,omitempty"`
	CreatedAt   time.Time        `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	DeletedAt   *time.Time       `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
	Digital     *DigitalDetails  `json:"digital,omitempty" bson:"digital,omitempty"`
	EndDate     *time.Time       `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ExternalIDs []string         `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	Filler      bool             `json:"filler,omitempty" bson:"filler,omitempty"`
	Lifecycle   BookingLifecycle `json:"lifecycle,omitempty" bson:"lifecycle,omitempty"`
	MarketID    string           `json:"marketID,omitempty" bson:"marketID,omitempty"`
	OrderID     string           `json:"orderID,omitempty" bson:"orderID,omitempty"`
	Print       *PrintDetails    `json:"print,omitempty" bson:"print,omitempty"`
	Product     ProductDetails   `json:"product,omitempty" bson:"product,omitempty"`
	Segment     SegmentDetails   `json:"segment,omitempty" bson:"segment,omitempty"`
	StartDate   *time.Time       `json:"startDate,omitempty" bson:"startDate,omitempty"`
	Status      BookingStatus    `json:"status,omitempty" bson:"status,omitempty"`
	UpdatedAt   time.Time        `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Waitlisted  *bool            `json:"waitlisted,omitempty" bson:"waitlisted,omitempty"`
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

type ProductDetails struct {
	ProductID    string `json:"productID,omitempty" bson:"productID,omitempty"`
	MediaProduct struct {
		ProductCode string `json:"productCode,omitempty" bson:"productCode,omitempty"`
		TypeCode    string `json:"typeCode,omitempty" bson:"typeCode,omitempty"`
	} `json:"media,omitempty" bson:"media,omitempty"`
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

func (b Booking) MarshalZerologObject(e *zerolog.Event) {
	canceled := ""
	if b.Canceled != nil {
		canceled = fmt.Sprintf("%v", *b.Canceled)
	}
	var start, end string
	if b.StartDate != nil {
		start = (*b.StartDate).Format(time.DateOnly)
	}
	if b.EndDate != nil {
		end = (*b.EndDate).Format(time.DateOnly)
	}
	waitlisted := ""
	if b.Waitlisted != nil {
		waitlisted = fmt.Sprintf("%v", *b.Waitlisted)
	}

	e.
		Str("model", "Booking").
		Str("bookingID", b.ID).
		Str("orderID", b.OrderID).
		Str("startDate", start).
		Str("endDate", end).
		Str("canceled", canceled).
		Bool("filler", b.Filler).
		Str("externalIDs", strings.Join(b.ExternalIDs, ", ")).
		Str("waitlisted", waitlisted)

	if b.Print != nil {
		e.Object("print", b.Print)
	}

	if b.Digital != nil {
		e.Object("digital", b.Digital)
	}
}

func (dd *DigitalDetails) MarshalZerologObject(e *zerolog.Event) {
	if dd == nil {
		return
	}

	e.
		Str("networkID", dd.NetworkID).
		Str("dailyStartTime", dd.DailyStartTime).
		Str("dailyEndTime", dd.DailyEndTime).
		Object("daysToPlay", dd.DaysToPlay).
		Int("numberOfSlots", dd.NumberOfSlots).
		Float32("slotSeconds", dd.SlotSeconds).
		Int("slotSlices", dd.SlotSlices).
		Str("specificStartTime", dd.SpecificStartTime)
}

func (pd *PrintDetails) MarshalZerologObject(e *zerolog.Event) {
	if pd == nil {
		return
	}

	e.Str("displayID", pd.DisplayID)
}

func (dtp *DigitalDaysToPlay) MarshalZerologObject(e *zerolog.Event) {
	if dtp == nil {
		return
	}

	e.
		Bool("sunday", dtp.Sunday).
		Bool("monday", dtp.Monday).
		Bool("tuesday", dtp.Tuesday).
		Bool("wednesday", dtp.Wednesday).
		Bool("thursday", dtp.Thursday).
		Bool("friday", dtp.Friday).
		Bool("saturday", dtp.Saturday)
}
