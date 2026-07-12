package domain

import "github.com/google/uuid"

type Shopping_cart_item struct {
	Item_id    uuid.UUID `json:"item_id" validate:"required"`
	Quantity   int       `json:"quantity" validate:"required,gt=0"`
	Unit_price float64   `json:"unit_price" validate:"required,gt=0"`
	Item_title *string   `json:"item_title" validate:"required"`
	Item_note  *string   `json:"item_note"`
}
