package commands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type Delete_catalog_item_handler struct {
	repo repositories.Catalog_item_repo
}

func New_delete_catalog_item_handler(repo repositories.Catalog_item_repo) *Delete_catalog_item_handler {
	return &Delete_catalog_item_handler{repo: repo}
}
func (h *Delete_catalog_item_handler) Handle(ctx context.Context,
	id uuid.UUID) (bool, error) {
	exist, err := h.repo.Item(ctx, id)
	if err != nil {
		return false, err
	}
	if exist == nil {
		return false, nil
	}

	deleted, err := h.repo.Delete(ctx, id)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
