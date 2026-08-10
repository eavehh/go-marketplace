package commands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/promotion/application/interfaces"
)

type Delete_promo_command struct {
	Catalog_item_id string
}
type Delete_promo_result struct {
	Success     bool
	Description string
}

type Delete_promo_handler struct {
	repo interfaces.Promorion_repository
}

func New_delete_promo_handler(repo interfaces.Promorion_repository) *Delete_promo_handler {
	return &Delete_promo_handler{repo: repo}
}

func (h *Delete_promo_handler) Handle(ctx context.Context, cmd Delete_promo_command) (*Delete_promo_result, error) {
	success, err := h.repo.Delete(ctx, cmd.Catalog_item_id)
	if err != nil {
		return nil, err
	}

	desc := "promo deleted successfully"
	if !success {
		desc = "promo not found or not deleted"
	}
	return &Delete_promo_result{
		Success:     success,
		Description: desc,
	}, nil

}
