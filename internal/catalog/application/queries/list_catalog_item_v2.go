package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/spec"
)

type Catalog_items_v2_handler struct {
	repo repositories.Catalog_item_repo
}

func New_catalog_items_v2_handler(repo repositories.Catalog_item_repo) *Catalog_items_v2_handler {
	return &Catalog_items_v2_handler{repo: repo}
}

func (q *Catalog_items_v2_handler) Handle(ctx context.Context, args spec.Query_args) (
	spec.Pagination[entity.Catalog_item], error) {
	return q.repo.Catalog_items(ctx, args)
}
