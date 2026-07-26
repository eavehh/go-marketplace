package persistence

import (
	"context"
	"database/sql"

	"github.com/eavehh/marketpl.microserv/internal/promotion/domain"
)

type Promo_repository struct {
	db *sql.DB
}

func New_promo_repository(db *sql.DB) *Promo_repository {
	return &Promo_repository{db: db}
}

func (r *Promo_repository) Find_by_catalog_item(ctx context.Context, catalog_item_id string) (*domain.Promo, error) {
	const query = `
	SELECT 
		id,
		catalog_item_id,
		title,
		value
	FROM promos
	WHERE catalog_item_id = ?
	LIMIT 1
	`
	var p domain.Promo

	err := r.db.QueryRowContext(ctx, query, catalog_item_id).Scan(
		&p.Id,
		&p.Catalog_item_id,
		&p.Title,
		&p.Value,
	)
 
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}
