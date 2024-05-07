package orders

type OrderSource string

const (
	Quattro     OrderSource = "Quattro"
	Hold        OrderSource = "Salesforce Hold"
	Opportunity OrderSource = "Salesforce Opportunity"
)

type SubmissionCategory string

const (
	OperationsIssue              SubmissionCategory = "Operations Issue"
	PrepayInvoiceRequest         SubmissionCategory = "Prepay Invoice Request"
	ProductionContract           SubmissionCategory = "Production Contract"
	InvoiceChange                SubmissionCategory = "Invoice Change"
	Cancellation                 SubmissionCategory = "Cancellation"
	BillingErrorCorrection       SubmissionCategory = "Billing Error Correction"
	HoldRequest                  SubmissionCategory = "Hold Request"
	PrepayInvoiceContractRequest SubmissionCategory = "Prepayment Invoice Contract Request"
	OrderBooking                 SubmissionCategory = "Order Booking"
	TradeOrderBooking            SubmissionCategory = "Trade Order Booking"
	PSAOrderBooking              SubmissionCategory = "PSA Order Booking"
	ComissionChange              SubmissionCategory = "Commission Change"
	RateChange                   SubmissionCategory = "Rate Change"
	ContractRevision             SubmissionCategory = "Contract Revision"
	BillingInquiry               SubmissionCategory = "Billing Inquiry"
)
