package handlers

import (
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type Catalog_items_handler struct {
	catalog_items *queries.Item_query_handler
}

func New_catalog_items_handler(Items *queries.Item_query_handler) *Catalog_items_handler {
	return &Catalog_items_handler{catalog_items: Items}
}

func (h *Catalog_items_handler) Items(c *gin.Context) {
	items, err := h.catalog_items.Handle(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(
		http.StatusOK, gin.H{
			"catalog_items": items,
		})
}
