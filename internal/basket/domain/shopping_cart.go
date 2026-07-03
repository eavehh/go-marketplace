package domain

type Shopping_cart struct {
	Account_name string               `json:"account_name"`
	Items        []Shopping_cart_item `json:"items"`
}

func (sc *Shopping_cart) Total_price() float64 {
	var total float64
	for _, item := range sc.Items {
		total += float64(item.Quantity) * item.Unit_price
	}
	return total
}
