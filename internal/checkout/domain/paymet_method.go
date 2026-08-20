package domain

type Payment_method string

const (
	Payment_method_credit_card   Payment_method = "Credit_card"
	Payment_method_bank_transfer Payment_method = "Bank_transfer"
)

func (m Payment_method) String() string {
	return string(m)
}

func (m Payment_method) Is_valid() bool {
	switch m {
	case Payment_method_credit_card,
		Payment_method_bank_transfer:
		return true
	}
	return false
}
