package commands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/promotion/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/promotion/domain"
)

type Update_promo_command struct {
	Id    string
	Title string
	Value string
}

type Update_promo_result struct {
	Success     bool
	Description string
}

type Update_promo_handler struct {
	repo interfaces.Promorion_repository
}

func New_update_promo_handler(repo interfaces.Promorion_repository) *Update_promo_handler {
	return &Update_promo_handler{repo: repo}
}

func (h *Update_promo_handler) Handle(ctx context.Context, cmd Update_promo_command) (*Update_promo_result, error) {
	promo := &domain.Promo{
		Id:    cmd.Id,
		Title: cmd.Title,
		Value: cmd.Value,
	}

	success, err := h.repo.Update(ctx, promo)
	if err != nil {
		return nil, err
	}

	desc := "updated"
	if !success {
		desc = "update error"
	}
	return &Update_promo_result{
		Success:     success,
		Description: desc,
	}, nil

}
