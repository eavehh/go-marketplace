package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Catalog_item_by_title_handler struct {
	repo repositories.Catalog_item_repo
}

func New_catalog_item_by_title_handler(repo repositories.Catalog_item_repo) *Catalog_item_by_title_handler {
	return &Catalog_item_by_title_handler{repo: repo}
}

func (q *Catalog_item_by_title_handler) Handle(ctx context.Context, title string) (
	[]entity.Catalog_item, error) {
	return q.repo.Item_by_title(ctx, title)
}
