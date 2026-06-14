package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Item_query_handler struct {
	repo repositories.Catalog_item_repo
}

func New_Item_queries(repo repositories.Catalog_item_repo) *Item_query_handler {
	return &Item_query_handler{repo: repo}
}

func (q *Item_query_handler) Handle(ctx context.Context) (
	[]entity.Catalog_item, error) {
	return q.repo.Items(ctx)
}
