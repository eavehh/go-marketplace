package domain

type Payment_status string

const (
	Payment_status_pending   Payment_status = "Pending"
	Payment_status_complited Payment_status = "Complited"
	Payment_status_faild     Payment_status = "Faild"
	Payment_status_refundend Payment_status = "Refundend"
)

func (m Payment_status) String() string {
	return string(m)
}

func (m Payment_status) Is_valid() bool {
	switch m {
	case Payment_status_pending,
		Payment_status_complited,
		Payment_status_faild,
		Payment_status_refundend:
		return true
	}
	return false
}
