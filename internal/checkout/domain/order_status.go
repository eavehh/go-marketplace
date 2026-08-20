package domain

type Order_status string

const (
	Order_status_draft     Order_status = "Draft"
	Order_status_submitted Order_status = "Submitted"
	Order_status_paid      Order_status = "Paid"
	Order_status_shipped   Order_status = "Shipped"
	Order_status_cancelled Order_status = "Cancelled"
)

func (s Order_status) String() string {
	return string(s)
}

func (s Order_status) Is_valid() bool {
	switch s {
	case Order_status_draft,
		Order_status_submitted,
		Order_status_paid,
		Order_status_shipped,
		Order_status_cancelled:
		return true
	}
	return false
}
