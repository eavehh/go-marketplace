package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Brands_query_handler struct {
	repo repositories.Brand_repo
}

func New_brands_queries(repo repositories.Brand_repo) *Brands_query_handler {
	return &Brands_query_handler{repo: repo}
}

func (q *Brands_query_handler) Handle(ctx context.Context) ([]entity.Brand, error) {
	return q.repo.Brands(ctx)
}
