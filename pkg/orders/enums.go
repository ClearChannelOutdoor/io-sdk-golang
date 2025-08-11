package orders

type AdditionalCostName string

const (
	ProductionDelivery AdditionalCostName = "Production Delivery"
	Installation       AdditionalCostName = "Installation"
	Embellishment      AdditionalCostName = "Embellishment"
	Illumniation       AdditionalCostName = "Illumination"
	CreativeDesign     AdditionalCostName = "Creative Design"
	ReturnShipping     AdditionalCostName = "Return Shipping"
)

type OrderSource string

const (
	Hold        OrderSource = "Salesforce Hold"
	Opportunity OrderSource = "Salesforce Opportunity"
	Quattro     OrderSource = "Quattro"
)

type SubmissionCategory string

const (
	BillingErrorCorrection       SubmissionCategory = "Billing Error Correction"
	BillingInquiry               SubmissionCategory = "Billing Inquiry"
	Cancellation                 SubmissionCategory = "Cancellation"
	ComissionChange              SubmissionCategory = "Commission Change"
	ContractRevision             SubmissionCategory = "Contract Revision"
	HoldRequest                  SubmissionCategory = "Hold Request"
	InvoiceChange                SubmissionCategory = "Invoice Change"
	OperationsIssue              SubmissionCategory = "Operations Issue"
	PrepayInvoiceRequest         SubmissionCategory = "Prepay Invoice Request"
	ProductionContract           SubmissionCategory = "Production Contract"
	PrepayInvoiceContractRequest SubmissionCategory = "Prepayment Invoice Contract Request"
	PSAOrderBooking              SubmissionCategory = "PSA Order Booking"
	OrderBooking                 SubmissionCategory = "Order Booking"
	RateChange                   SubmissionCategory = "Rate Change"
	TradeOrderBooking            SubmissionCategory = "Trade Order Booking"
)
