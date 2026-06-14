package repositories

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/google/uuid"
)

type Catalog_item_repo interface {
	Items(ctx context.Context) ([]entity.Catalog_item, error)
	Item(ctx context.Context, id uuid.UUID) (*entity.Catalog_item, error)
}

type Brand_repo interface {
	Brands(ctx context.Context) ([]entity.Brand, error)
}

type Category_repo interface {
	Categories(ctx context.Context) ([]entity.Category, error)
}
