package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/checkout/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/checkout/domain"
)

type Orders_by_account_name_handler struct {
	repo interfaces.Order_repository
}
type Orders_by_account_name struct {
	Account_name string
}

func New_order_by_account_name_handler(repo interfaces.Order_repository) *Orders_by_account_name_handler {
	return &Orders_by_account_name_handler{repo: repo}
}

func (h *Orders_by_account_name_handler) Handle(ctx context.Context, q Orders_by_account_name) ([]domain.Order, error) {
	return h.repo.Find_by_account_name(ctx, q.Account_name)
}
