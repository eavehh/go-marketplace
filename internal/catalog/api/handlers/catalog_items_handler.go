package handlers

import (
	"log"
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Catalog_items_handler struct {
	catalog_items *queries.Catalog_items_handler
	item_by_id    *queries.Catalog_item_by_id_handler
}

func New_catalog_items_handler(
	items *queries.Catalog_items_handler,
	item_by_id *queries.Catalog_item_by_id_handler,
) *Catalog_items_handler {
	return &Catalog_items_handler{
		catalog_items: items,
		item_by_id:    item_by_id,
	}
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

func (h *Catalog_items_handler) Item_by_id(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		log.Println("file catalog_item_handler:", err)
		return
	}

	item, err := h.item_by_id.Handle(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if item == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "catalog item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": item,
	})
}
