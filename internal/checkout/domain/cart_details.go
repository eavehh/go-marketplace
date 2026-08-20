package domain

type Cart_details struct {
	Card_name   string `json:"card_name" validate:"max=100"`
	Cart_number string `json:"cart_number" validate:"max=20"`
	Expiration  string `json:"expiration" validate:"max=10"`
	CVV         string `json:"cvv" validate:"max=10"`
}
