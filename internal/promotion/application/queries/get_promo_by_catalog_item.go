package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/promotion/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/promotion/domain"
)

type Get_by_catalog_item_query struct {
	Catalog_item_id string
}

type Get_by_catalog_item_handler struct {
	repo interfaces.Promorion_repository
}

func New_get_by_catalog_item_handler(repo interfaces.Promorion_repository) *Get_by_catalog_item_handler {
	return &Get_by_catalog_item_handler{repo: repo}
}

func (h *Get_by_catalog_item_handler) Handle(ctx context.Context, q Get_by_catalog_item_query) (*domain.Promo, error) {
	return h.repo.Find_by_catalog_item(ctx, q.Catalog_item_id)
}
