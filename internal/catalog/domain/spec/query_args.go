package spec

import "github.com/google/uuid"

const Max_page_size = 4

type Query_args struct {
	Page_index  int     `form"page_index"`
	Page_size   int     `form"page_size"`
	Brand_id    *string `form"brand_id"`
	Category_id *string `form"category_id"`
	Search      *string `form"search"`
	Sort        *string `form"sort"`
}

func (q *Query_args) Normalize() {
	if q.Page_index < 1 {
		q.Page_index = 1
	}
	if q.Page_size <= 0 {
		q.Page_size = 2
	}
	if q.Page_size > Max_page_size {
		q.Page_size = Max_page_size
	}
}

func (q *Query_args) Parse_brand_id() (*uuid.UUID, error) {
	if q.Brand_id == nil || *q.Brand_id == "" {
		return nil, nil
	}

	id, err := uuid.Parse(*q.Brand_id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (q *Query_args) Parse_category_id() (*uuid.UUID, error) {
	if q.Category_id == nil || *q.Category_id == "" {
		return nil, nil
	}
	id, err := uuid.Parse(*q.Category_id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
