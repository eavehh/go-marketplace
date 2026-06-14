package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type Catalog_item_by_id_handler struct {
	repo repositories.Catalog_item_repo
}

func New_catalog_item_by_id_handler(repo repositories.Catalog_item_repo) *Catalog_item_by_id_handler {
	return &Catalog_item_by_id_handler{repo: repo}
}

func (q *Catalog_item_by_id_handler) Handle(ctx context.Context, id uuid.UUID) (
	*entity.Catalog_item, error) {
	return q.repo.Item(ctx, id)
}
