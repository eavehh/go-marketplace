package repositories

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
)

type Catalog_item_repo interface {
	Item(ctx context.Context) ([]entity.Catalog_item, error)
}

type Brand_repo interface {
	Brand(ctx context.Context) ([]entity.Brand, error)
}

type Category_repo interface {
	Category(ctx context.Context) ([]entity.Category, error)
}
