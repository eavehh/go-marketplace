package persistence

import (
	"context"
	"database/sql"

	"github.com/eavehh/marketpl.microserv/internal/checkout/domain"
	"github.com/google/uuid"
)

type Order_repository struct {
	db *sql.DB
}

func New_order_repository(db *sql.DB) *Order_repository {
	return &Order_repository{db: db}
}

func (r *Order_repository) Get(ctx context.Context, id uuid.UUID) (*domain.Order, error) {

	const query = `
	SELECT id, account_name, total_amount,
	current_order_status,
	contact_first_name, contact_last_name, contact_email, address_street,
	address_city, address_region, address_postal_code, current_payment_method,
	current_payment_status, card_name, card_number, card_expiration, card_cvv,
	created_by, created_at, last_modified_by, last_modified_at
	FROM orders
	WHERE id = $1
	`
	order := domain.Order{}
	var card_name, card_number, card_expiration, card_cvv sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&order.Id,
		&order.Account_name,
		&order.Total_amount,
		&order.Current_order_status,
		&order.Contact_info.First_name,
		&order.Contact_info.Last_name,
		&order.Contact_info.Email,
		&order.Delivery_address.Street,
		&order.Delivery_address.City,
		&order.Delivery_address.Region,
		&order.Delivery_address.Postal_code,
		&order.Current_payment_method,
		&order.Current_payment_status,
		&card_name,
		&card_number,
		&card_expiration,
		&card_cvv,
		&order.Created_by,
		&order.Created_at,
		&order.Last_modified_by,
		&order.Last_modified_at,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if card_number.Valid {
		order.Cart_details = &domain.Cart_details{
			Cart_number: card_number.String,
			Card_name:   card_name.String,
			Expiration:  card_expiration.String,
			CVV:         card_cvv.String,
		}
	}

	items, err := r.get_order_items(ctx, id)
	if err != nil {
		return nil, err
	}

	order.Items = items
	return &order, nil
}

func (r *Order_repository) get_order_items(ctx context.Context, order_id uuid.UUID) ([]domain.Order_item, error) {
	const query = `
		SELECT
			catalog_item_name,
			quantity,
			unit_price
		FROM order_items
		WHERE order_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, order_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Order_item

	for rows.Next() {
		var item domain.Order_item
		err := rows.Scan(
			&item.Catalog_item_name,
			&item.Quantity,
			&item.Unit_price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *Order_repository) Find_by_account_name(ctx context.Context, account_name string) ([]domain.Order, error) {
	return nil, nil

}

func (r *Order_repository) Create(ctx context.Context, order *domain.Order) error {
	return nil
}
