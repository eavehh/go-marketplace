package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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

func (r *Promo_repository) Create(ctx context.Context, promo *domain.Promo) (bool, error) {
	query :=
		`
	INSERT INTO promos(id, catalog_item_id, title, value)
	VALUES (?,?,?,?)
	`

	result, err := r.db.ExecContext(ctx, query,
		promo.Id,
		promo.Catalog_item_id,
		promo.Title,
		promo.Value,
	)

	if err != nil {

		if (strings.Contains(err.Error(), "Dublicate entity")) ||
			strings.Contains(err.Error(), "unique_catalog_item_id") {
			return false, fmt.Errorf("promo for catalog_item_id (%s) alredy exist: %v", promo.Catalog_item_id, err)
		}
		return false, fmt.Errorf("repo error: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rows > 0, nil
}

func (r *Promo_repository) Update(ctx context.Context, promo *domain.Promo) (bool, error) {
	query := `
	 	UPDATE promos SET title = ?, value = ?
		WHERE id = ?
	 `

	result, err := r.db.ExecContext(ctx, query, promo.Title, promo.Value, promo.Id)

	if err != nil {
		return false, err
	}

	rows, _ := result.RowsAffected()
	return rows > 0, nil

}

func (r *Promo_repository) Delete(ctx context.Context, catalog_item_id string) (bool, error) {
	query := `DELETE FROM promos WHERE id = $1`

	result, err := r.db.Exec(query, catalog_item_id)
	if err != nil {
		return false, fmt.Errorf("Delete promo: %v, error: %w", catalog_item_id, err)
	}

	n, _ := result.RowsAffected()
	return n > 0, nil

}
