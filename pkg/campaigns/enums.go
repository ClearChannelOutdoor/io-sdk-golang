package campaigns

type Deliverable string

const (
	Display     Deliverable = "Display"
	Impressions Deliverable = "Impressions"
	Quantity    Deliverable = "Quantity"
)

type Flexibility string

const (
	Fixed    Flexibility = "Fixed"
	Flexible Flexibility = "Flexible"
)
