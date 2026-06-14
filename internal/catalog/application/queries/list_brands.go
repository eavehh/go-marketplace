package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Brands_handler struct {
	repo repositories.Brand_repo
}

func New_brands_handler(repo repositories.Brand_repo) *Brands_handler {
	return &Brands_handler{repo: repo}
}

func (q *Brands_handler) Handle(ctx context.Context) ([]entity.Brand, error) {
	return q.repo.Brands(ctx)
}
