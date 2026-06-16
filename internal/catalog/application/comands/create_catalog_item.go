package comands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type Create_catalog_item_command struct {
	Title             string
	Short_description string
	Full_description  string
	Image_url         string
	Price             float64
	Brand             entity.Brand
	Category          entity.Category
}

type Create_catalog_item_handler struct {
	repo repositories.Catalog_item_repo
}

func New_create_catalog_item_handler(repo repositories.Catalog_item_repo) *Create_catalog_item_handler {
	return &Create_catalog_item_handler{repo: repo}
}
func (h *Create_catalog_item_handler) Handle(ctx context.Context,
	cmd Create_catalog_item_command) (uuid.UUID, error) {
	item := entity.Catalog_item{
		Base_entity: entity.Base_entity{
			Id:    uuid.New(),
			Title: cmd.Title,
		},
		Short_description: &cmd.Short_description,
		Full_description:  &cmd.Full_description,
		Image_url:         &cmd.Image_url,
		Price:             cmd.Price,
		Brand:             &cmd.Brand,
		Category:          &cmd.Category,
	}

	created, err := h.repo.Create(ctx, item)
	if err != nil {
		return uuid.Nil, err
	}

	return created.Id, nil
}
