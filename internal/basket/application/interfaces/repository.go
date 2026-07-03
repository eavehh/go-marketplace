package interfaces

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
)

type Cart_repository interface {
	Save(ctx context.Context, shop_cart_item *domain.Shopping_cart) (*domain.Shopping_cart, error)
}
