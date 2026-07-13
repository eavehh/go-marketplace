package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/basket/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
)

type Get_cart_handler struct {
	repo interfaces.Cart_repository
}

func New_get_cart_handler(repo interfaces.Cart_repository) *Get_cart_handler {
	return &Get_cart_handler{repo: repo}
}

func (h *Get_cart_handler) Handle(ctx context.Context, account_name string,
) (*domain.Shopping_cart, error) {
	return h.repo.Get(ctx, account_name)
}
