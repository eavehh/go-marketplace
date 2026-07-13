package interfaces

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
)

type Cart_repository interface {
	Save(ctx context.Context, shopping_cart_items *domain.Shopping_cart) (*domain.Shopping_cart, error)
	Get(ctx context.Context, account_name string) (*domain.Shopping_cart, error)
}
