package domain

type Order_item struct {
	Catalog_item_name string  `json:"catalog_item_name" validate:"required,max=200"`
	Quantity          int     `json:"quantity" validate:"required,gt=0"`
	Unit_price        float64 `json:"unit_price" validate:"required,gte=0"`
}

// return total price of the position
func (i Order_item) Total() float64 {
	return float64(i.Quantity) * i.Unit_price
}
