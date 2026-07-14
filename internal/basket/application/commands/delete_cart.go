package commands

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/basket/application/interfaces"
)

type Delete_cart_handler struct {
	repo interfaces.Cart_repository
}

func New_delete_handler(repo interfaces.Cart_repository) *Delete_cart_handler {
	return &Delete_cart_handler{repo: repo}
}

func (h *Delete_cart_handler) Handle(ctx context.Context, account_name string) error {
	return h.repo.Delete(ctx, account_name)
}
