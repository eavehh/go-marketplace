package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Categories_handler struct {
	repo repositories.Category_repo
}

func New_categories_handler(repo repositories.Category_repo) *Categories_handler {
	return &Categories_handler{repo: repo}
}
func (q *Categories_handler) Handle(ctx context.Context) ([]entity.Category, error) {
	return q.repo.Categories(ctx)
}
