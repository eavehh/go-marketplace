package domain

import "github.com/google/uuid"

type Shop_cart_item struct {
	Item_id    uuid.UUID `json:"item_id"`
	Quantity   int       `json:"quantity"`
	Unit_price float64   `json:"unit_price"`
	Item_title *string   `json:"item_title"`
	Item_note  *string   `json:"item_note"`
}
