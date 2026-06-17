package commands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type Update_catalog_item_command_model struct {
	Id                *uuid.UUID      `json:"id"`
	Title             string          `json:"title"`
	Short_description string          `json:"short_description"`
	Full_description  string          `json:"full_description"`
	Image_url         string          `json:"image_url"`
	Price             float64         `json:"price"`
	Brand             entity.Brand    `json:"brand"`
	Category          entity.Category `json:"category"`
}

type Update_catalog_item_handler struct {
	repo repositories.Catalog_item_repo
}

func New_ureate_catalog_item_handler(repo repositories.Catalog_item_repo) *Update_catalog_item_handler {
	return &Update_catalog_item_handler{repo: repo}
}
func (h *Update_catalog_item_handler) Handle(ctx context.Context,
	cmd Update_catalog_item_command_model) (bool, error) {
	exist, err := h.repo.Item(ctx, *cmd.Id)
	if err != nil {
		return false, err
	}
	if exist == nil {
		return false, nil
	}

	item := entity.Catalog_item{
		Base_entity: entity.Base_entity{
			Id:    *cmd.Id,
			Title: cmd.Title,
		},
		Short_description: &cmd.Short_description,
		Full_description:  &cmd.Full_description,
		Image_url:         &cmd.Image_url,
		Price:             cmd.Price,
		Brand:             &cmd.Brand,
		Category:          &cmd.Category,
	}

	updated, err := h.repo.Update(ctx, item)
	if err != nil {
		return updated, err
	}

	return updated, nil
}
