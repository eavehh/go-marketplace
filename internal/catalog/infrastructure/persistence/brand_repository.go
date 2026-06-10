package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
)

type brand_repo struct {
	db *sql.DB
}

func New_brand_repo(db *sql.DB) *brand_repo {
	return &brand_repo{db: db}
}

func (r *brand_repo) Brands(ctx context.Context) ([]entity.Brand, error) {
	rows, err := r.db.QueryContext(ctx, `
	SELECT id, title, image_url  FROM brands ORDER by title`)
	if err != nil {
		return nil, fmt.Errorf("brands query: %w", err)
	}

	defer rows.Close()

	var brands []entity.Brand

	for rows.Next() {
		var b entity.Brand
		if err := rows.Scan(&b.Id, &b.Title, &b.Image_url); err != nil {
			return nil, fmt.Errorf("scan brands %w", err)
		}
		brands = append(brands, b)
	}

	return brands, rows.Err()
}
