package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id               uuid.UUID  `json:"id"`
	Created_by       *string    `json:"created_by,omitempty"`
	Created_at       time.Time  `json:"created_at"`
	Last_modified_by *string    `json:"last_modified_by,omitempty"`
	Last_modified_at *time.Time `json:"last_modified_at,omitempty"`

	Account_name           string         `json:"account_name" validate:"required,max=100"`
	Total_amount           float64        `json:"total_amount"`
	Items                  []Order_item   `json:"items" validate:"required,min=1,dive"`
	Current_order_status   Order_status   `json:"current_order_status"`
	Contact_info           Contact        `json:"contact_info" validate:"required"`
	Delivery_address       Address        `json:"delivery_address" validate:"required"`
	Current_payment_method Payment_method `json:"current_payment_method"`
	Cart_details           *Cart_details  `json:"cart_details"`
	Current_payment_status Payment_status `json:"current_payment_status"`
}

func (o Order) New_order() *Order {
	return &Order{
		Id:                     uuid.New(),
		Created_at:             time.Now().UTC(),
		Current_order_status:   Order_status_draft,
		Current_payment_status: Payment_status_pending,
		Items:                  make([]Order_item, 0),
	}
}

func (o *Order) Set_created_audit(Created_by string) {
	o.Created_by = &Created_by
	o.Created_at = time.Now().UTC()
}

func (o *Order) Set_modifies_audit(modified_by string) {
	o.Last_modified_by = &modified_by
	last_modified_at := time.Now().UTC()
	o.Last_modified_at = &last_modified_at
}

func (o *Order) Calculate_total_amount() {
	var total float64

	for _, item := range o.Items {
		total += item.Total()
	}
	o.Total_amount = total
}
