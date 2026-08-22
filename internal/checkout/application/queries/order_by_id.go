package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/checkout/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/checkout/domain"
	"github.com/google/uuid"
)

type Order_by_id_query_handler struct {
	repo interfaces.Order_repository
}
type Order_by_id_query struct {
	Id uuid.UUID
}

func New_order_by_id_query_handler(repo interfaces.Order_repository) *Order_by_id_query_handler {
	return &Order_by_id_query_handler{repo: repo}
}

func (h *Order_by_id_query_handler) Handle(ctx context.Context, q Order_by_id_query) (*domain.Order, error) {
	return h.repo.Get(ctx, q.Id)
}
