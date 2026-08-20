package interfaces

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/checkout/domain"
	"github.com/google/uuid"
)

type Order_repository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Order, error)
	Find_by_account_name(ctx context.Context, account_name string) ([]domain.Order, error)
	Create(ctx context.Context, order *domain.Order) error
}
