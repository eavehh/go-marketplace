package queries

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/entity"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/repositories"
)

type Categories_query_handler struct {
	repo repositories.Category_repo
}

func New_categories_query(repo repositories.Category_repo) *Categories_query_handler {
	return &Categories_query_handler{repo: repo}
}
func (q *Categories_query_handler) Handle(ctx context.Context) ([]entity.Category, error) {
	return q.repo.Categories(ctx)
}
