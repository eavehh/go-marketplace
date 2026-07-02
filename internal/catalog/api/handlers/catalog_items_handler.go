package handlers

import (
	"log"
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/catalog/application/commands"
	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Catalog_items_handler struct {
	catalog_items       *queries.Catalog_items_handler
	item_by_id          *queries.Catalog_item_by_id_handler
	item_by_title       *queries.Catalog_item_by_title_handler
	item_by_brand_title *queries.Catalog_item_by_brand_title_handler

	create_item *commands.Create_catalog_item_handler
	update_item *commands.Update_catalog_item_handler
	delete_item *commands.Delete_catalog_item_handler
}

func New_catalog_items_handler(
	items *queries.Catalog_items_handler,
	item_by_id *queries.Catalog_item_by_id_handler,
	item_by_title *queries.Catalog_item_by_title_handler,
	item_by_brand_title *queries.Catalog_item_by_brand_title_handler,
	create_item *commands.Create_catalog_item_handler,
	update_item *commands.Update_catalog_item_handler,
	delete_item *commands.Delete_catalog_item_handler,

) *Catalog_items_handler {
	return &Catalog_items_handler{
		catalog_items:       items,
		item_by_id:          item_by_id,
		item_by_title:       item_by_title,
		item_by_brand_title: item_by_brand_title,

		create_item: create_item,
		update_item: update_item,
		delete_item: delete_item,
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
		c.JSON(http.StatusNotFound, gin.H{
			"error": "catalog item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": item,
	})
}

func (h *Catalog_items_handler) Item_by_title(c *gin.Context) {
	title := c.Param("title")

	items, err := h.item_by_title.Handle(c.Request.Context(), title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": items,
	})
}

func (h *Catalog_items_handler) Item_by_brand_title(c *gin.Context) {
	brand_title := c.Param("brand")

	items, err := h.item_by_brand_title.Handle(c, brand_title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": items,
	})

}

func (h *Catalog_items_handler) Create_item(c *gin.Context) {
	var cmd_model commands.Create_catalog_item_command_model
	if err := c.ShouldBindJSON(&cmd_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := h.create_item.Handle(c, cmd_model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *Catalog_items_handler) Update_item(c *gin.Context) {
	var cmd_model commands.Update_catalog_item_command_model
	if err := c.ShouldBindJSON(&cmd_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	updated, err := h.update_item.Handle(c, cmd_model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_success": updated,
	})

}

func (h *Catalog_items_handler) Delete_item(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	success, err := h.delete_item.Handle(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_success": success,
	})

}
