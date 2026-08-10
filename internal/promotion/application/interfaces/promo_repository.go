package interfaces

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/promotion/domain"
)

type Promorion_repository interface {
	Find_by_catalog_item(ctx context.Context, catalog_item_id string) (*domain.Promo, error)
	Create(ctx context.Context, promo *domain.Promo) (bool, error)
	Update(ctx context.Context, promo *domain.Promo) (bool, error)
	Delete(ctx context.Context, catalog_item_id string) (bool, error)
}
