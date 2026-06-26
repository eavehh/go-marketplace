package spec

type Pagination[T any] struct {
	Page_index  int `form:"page_index"`
	Page_size   int `form:"page_size"`
	Total_count int `form:"total_count"`
	Items       []T `form:"items"`
}
