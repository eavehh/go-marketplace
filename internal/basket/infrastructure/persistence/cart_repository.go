package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
	"github.com/google/uuid"
)

type Cart_repository struct {
	db *sql.DB
}

func New_cart_repository(db *sql.DB) *Cart_repository {
	return &Cart_repository{db: db}
}

func (r *Cart_repository) Save(ctx context.Context, cart *domain.Shopping_cart) (*domain.Shopping_cart, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`INSERT INTO shopping_carts (account_name)
	VALUES ($1)
	ON CONFLICT (account_name) DO NOTHING`,
		cart.Account_name,
	)

	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM shopping_cart_items
	WHERE account_name = $1`,
		cart.Account_name,
	)

	if err != nil {
		return nil, err
	}

	for i := range cart.Items {
		item := &cart.Items[i]
		if item.Item_id == uuid.Nil {
			item.Item_id = uuid.New()
		}

		_, err = tx.ExecContext(ctx,
			`INSERT INTO shopping_cart_items
			(account_name, item_id, quantity, unit_price, item_title, item_note)
			VALUES ($1, $2, $3, $4, $5, $6)`,
			cart.Account_name,
			item.Item_id,
			item.Quantity,
			item.Unit_price,
			item.Item_title,
			item.Item_note,
		)
		if err != nil {
			return nil, err
		}

	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *Cart_repository) Get(ctx context.Context, account_name string) (*domain.Shopping_cart, error) {
	var exist bool

	err := r.db.QueryRowContext(ctx,
		`SELECT EXIST(
	SELECT 1 FROM shopping_carts
	WHERE account_name = $1
	)`,
		account_name,
	).Scan(&exist)

	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("cart for %s does not exist", account_name)
	}

	rows, err := r.db.QueryContext(ctx,
		`SELECT item_id quantity unit_price item_title item_note
	FROM shopping_cart_items
	WHERE account_name = $1
	`,
		account_name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []domain.Shopping_cart_item
	for rows.Next() {
		var item domain.Shopping_cart_item
		err := rows.Scan(
			&item.Item_id,
			&item.Quantity,
			&item.Unit_price,
			&item.Item_title,
			&item.Item_note,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &domain.Shopping_cart{
		Account_name: account_name,
		Items:        items,
	}, nil
}
