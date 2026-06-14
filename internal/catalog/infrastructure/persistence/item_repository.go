package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/google/uuid"
)

type item_repository struct {
	db *sql.DB
}

func New_item_repository(db *sql.DB) *item_repository {
	return &item_repository{db: db}
}

func (r *item_repository) Items(ctx context.Context) ([]entity.Catalog_item, error) {
	query := `
	SELECT
	ci.id,
	ci.title,
	ci.short_description,     
	ci.full_description,
	ci.image_url,     
	ci.price,
	brnd.id,
	brnd.title,
	ctg.id,
	ctg.title
	FROM catalog_items ci
	LEFT JOIN brands brnd ON ci.brand_id = brnd.id
	LEFT JOIN categories ctg ON ci.category_id = ctg.id
	`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("items query: %w", err)
	}

	defer rows.Close()

	var items []entity.Catalog_item
	for rows.Next() {
		item, err := scan_ci(rows)

		if err != nil {
			return nil, fmt.Errorf("scan catalog item: %w", err)
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func scan_ci(rows *sql.Rows) (entity.Catalog_item, error) {
	var item entity.Catalog_item
	var brand_id *uuid.UUID
	var brand_title *string
	var category_id *uuid.UUID
	var category_title *string

	err := rows.Scan(
		&item.Id,
		&item.Title,
		&item.Short_description,
		&item.Full_description,
		&item.Image_url,
		&item.Price,
		&brand_id,
		&brand_title,
		&category_id,
		&category_title,
	)

	if err != nil {
		return item, err
	}

	if brand_id != nil {
		item.Brand = &entity.Brand{
			Base_entity: entity.Base_entity{
				Id:    *brand_id,
				Title: *brand_title,
			},
		}
	}

	if category_id != nil {
		item.Category = &entity.Category{
			Base_entity: entity.Base_entity{
				Id:    *category_id,
				Title: *category_title,
			},
		}
	}

	return item, nil
}
