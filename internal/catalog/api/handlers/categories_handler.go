package handlers

import (
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type Categories_handler struct {
	categories *queries.Categories_query_handler
}

func New_categories_handler(categories *queries.Categories_query_handler) *Categories_handler {
	return &Categories_handler{categories: categories}
}

func (h *Categories_handler) Categories(c *gin.Context) {
	categories, err := h.categories.Handle(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(
		http.StatusOK, gin.H{
			"categories": categories,
		})
}
