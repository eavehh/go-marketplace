package handlers

import (
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/eavehh/marketpl.microserv/internal/catalog/domain/spec"
	"github.com/gin-gonic/gin"
)

type Catalog_items_handler_v2 struct {
	Catalog_items_v2_handler *queries.Catalog_items_v2_handler
}

func New_catalog_items_handler_v2(query_handler *queries.Catalog_items_v2_handler) *Catalog_items_handler_v2 {
	return &Catalog_items_handler_v2{Catalog_items_v2_handler: query_handler}
}

func (h *Catalog_items_handler_v2) Catalog_items(c *gin.Context) {
	var args spec.Query_args
	err := c.ShouldBindQuery(&args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	args.Normalize()
	items, err := h.Catalog_items_v2_handler.Handle(c.Request.Context(), args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"catalog_items": items,
	})
}
