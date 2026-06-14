package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Catalog_items_handler struct {
	repo repositories.Catalog_item_repo
}

func New_catalog_items_handler(repo repositories.Catalog_item_repo) *Catalog_items_handler {
	return &Catalog_items_handler{repo: repo}
}

func (q *Catalog_items_handler) Handle(ctx context.Context) (
	[]entity.Catalog_item, error) {
	return q.repo.Items(ctx)
}
