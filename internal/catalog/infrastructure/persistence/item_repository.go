package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/google/uuid"
)

const sql_catalog_items_query = `
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

type item_repository struct {
	db *sql.DB
}

func New_item_repository(db *sql.DB) *item_repository {
	return &item_repository{db: db}
}

func (r *item_repository) Items(ctx context.Context) ([]entity.Catalog_item, error) {
	query := sql_catalog_items_query

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("items query: %w", err)
	}

	defer rows.Close()

	var items []entity.Catalog_item
	for rows.Next() {
		item, err := scan_catalog_item(rows)

		if err != nil {
			return nil, fmt.Errorf("scan catalog item: %w", err)
		}
		items = append(items, *item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("items iteration: %w", err)
	}

	return items, nil
}

func (r *item_repository) Item(ctx context.Context, id uuid.UUID) (*entity.Catalog_item, error) {
	query := sql_catalog_items_query + " WHERE ci.id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	item, err := scan_catalog_item(row)

	// errors.Is решает проблему обертки внутри err из-за которой err не может равняться sql.ErrNoRows (я завернул ее в фугкции scan_catalog_item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return item, nil
}

func (r *item_repository) Item_by_title(ctx context.Context, title string) ([]entity.Catalog_item, error) {
	query := sql_catalog_items_query + `WHERE ci.title ILIKE '%' || $1 || '%'`

	rows, err := r.db.QueryContext(ctx, query, title)

	if err != nil {
		return nil, fmt.Errorf("item by title query: %w", err)
	}

	defer rows.Close()

	var items []entity.Catalog_item = []entity.Catalog_item{}
	for rows.Next() {
		item, err := scan_catalog_item(rows)

		if err != nil {
			return nil, fmt.Errorf("scan catalog item by title: %w", err)
		}
		items = append(items, *item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("items by title iteration: %w", err)
	}
	// я бы добавил логирование включая название файла или класса
	return items, nil
}

func (r *item_repository) Create(ctx context.Context, item entity.Catalog_item) (entity.Catalog_item, error) {
	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}

	var brand_id, category_id uuid.UUID
	if item.Brand != nil {
		brand_id = item.Brand.Id
	}
	if item.Category != nil {
		category_id = item.Category.Id
	}

	sql_create_query := `
	INSERT INTO catalog_items(
    id, 
    title,
    short_description,
    full_description,
    image_url,
    price,
    brand_id,
    category_id
    )
	VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := r.db.Exec(sql_create_query,
		item.Id,
		item.Title,
		item.Short_description,
		item.Full_description,
		item.Image_url,
		item.Price,
		brand_id,
		category_id,
	)

	if err != nil {
		return entity.Catalog_item{}, fmt.Errorf("create item: %w", err)
	}

	return item, nil
}

func (r *item_repository) Update(ctx context.Context, item entity.Catalog_item) (bool, error) {
	var brand_id, category_id *uuid.UUID

	if item.Brand != nil {
		brand_id = &item.Brand.Id
	}
	if item.Category != nil {
		category_id = &item.Category.Id
	}

	sql_update_query := `
	UPDATE catalog_items SET
    title = $1,
    short_description = $2,
    full_description = $3,
    image_url = $4,
    price = $5,
    brand_id = $6,
    category_id = $7
	WHERE id = $8`

	result, err := r.db.ExecContext(ctx, sql_update_query,
		item.Title,
		item.Short_description,
		item.Full_description,
		item.Image_url,
		item.Price,
		brand_id,
		category_id,
		item.Id,
	)

	if err != nil {
		return false, fmt.Errorf("update item: item[%v], error: %w", item.Id, err)
	}

	n, _ := result.RowsAffected()
	return n > 0, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scan_catalog_item(scanner_rows scanner) (*entity.Catalog_item, error) {
	var item *entity.Catalog_item = &entity.Catalog_item{}
	var brand_id *uuid.UUID
	var brand_title *string
	var category_id *uuid.UUID
	var category_title *string

	err := scanner_rows.Scan(
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
		return item, fmt.Errorf("scan catalog item: %w", err) // будет продлема со сравнением ошилки из-заобертки
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
