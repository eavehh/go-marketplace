package commands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/basket/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Save_cart_command struct {
	Cart domain.Shopping_cart `json:"cart" validate:"required" `
}

type Save_cart_handler struct {
	repo interfaces.Cart_repository
}

func New_save_cart_handler(repo interfaces.Cart_repository) *Save_cart_handler {
	return &Save_cart_handler{repo: repo}
}

func (h *Save_cart_handler) Handle(ctx context.Context, cmd Save_cart_command) (string, error) {
	if err := validate.Struct(&cmd); err != nil {
		return "", err
	}

	_, err := h.repo.Save(ctx, &cmd.Cart)
	if err != nil {
		return "", err
	}
	return cmd.Cart.Account_name, nil
}
