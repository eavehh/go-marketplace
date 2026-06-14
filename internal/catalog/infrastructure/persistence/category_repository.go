package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
)

type category_repo struct {
	db *sql.DB
}

func New_Category_repo(db *sql.DB) *category_repo {
	return &category_repo{db: db}
}

func (r *category_repo) Categories(ctx context.Context) ([]entity.Category, error) {
	rows, err := r.db.QueryContext(ctx, `
	SELECT id, title FROM categories ORDER by title`)
	if err != nil {
		return nil, fmt.Errorf("categories query: %w", err)
	}

	defer rows.Close()

	var categories []entity.Category

	for rows.Next() {
		var c entity.Category
		if err := rows.Scan(&c.Id, &c.Title); err != nil {
			return nil, fmt.Errorf("scan categories: %w", err)
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}
