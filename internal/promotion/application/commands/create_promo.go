package commands

import (
	"context"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/promotion/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/promotion/domain"
	"github.com/google/uuid"
)

type Create_promo_command struct {
	Catalog_item_id string
	Title           string
	Value           string
}
type Create_promo_result struct {
	Id          string
	Success     bool
	Description string
}

type Create_promo_handler struct {
	repo interfaces.Promorion_repository
}

func New_create_promo_handler(repo interfaces.Promorion_repository) *Create_promo_handler {
	return &Create_promo_handler{repo: repo}
}

func (h *Create_promo_handler) Handle(ctx context.Context, cmd Create_promo_command) (*Create_promo_result, error) {
	exist_promo, err := h.repo.Find_by_catalog_item(ctx, cmd.Catalog_item_id)
	if err != nil {
		return &Create_promo_result{
			Id:          "",
			Success:     false,
			Description: "error check existing promo" + err.Error(),
		}, nil
	}

	if exist_promo != nil {
		msg := fmt.Sprintf("Promo already exist for %s ID (%s)",
			cmd.Catalog_item_id,
			exist_promo.Id,
		)
		return &Create_promo_result{
			Id:          exist_promo.Id,
			Success:     false,
			Description: msg,
		}, nil
	}

	promo := &domain.Promo{
		Id:              uuid.New().String(),
		Catalog_item_id: cmd.Catalog_item_id,
		Title:           cmd.Title,
		Value:           cmd.Value,
	}

	success, err := h.repo.Create(ctx, promo)

	if err != nil {
		return &Create_promo_result{
			Id:          promo.Id,
			Success:     false,
			Description: "error creating promo " + err.Error(),
		}, nil
	}

	description := "error creating promo"

	if success {
		description = "promo created successfully"
	}
	return &Create_promo_result{
		Id:          promo.Id,
		Success:     success,
		Description: description,
	}, nil
}
