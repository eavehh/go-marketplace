package api

import (
	"github.com/eavehh/marketpl.microserv/internal/catalog/api/handlers"
	"github.com/gin-gonic/gin"
)

func Register_routes(r *gin.Engine,
	brands *handlers.Brands_handler,
	categories *handlers.Categories_handler,
	items *handlers.Catalog_items_handler,
	items_v2 *handlers.Catalog_items_handler_v2,
) {
	v1 := r.Group("/api/v1")

	v1.GET("/brands", brands.Brands)
	v1.GET("/categories", categories.Categories)
	v1.GET("/catalog-items", items.Items)
	v1.GET("/catalog-items/:id", items.Item_by_id)
	v1.GET("/catalog-items/title/:title", items.Item_by_title)
	v1.GET("/catalog-items/brand/:brand", items.Item_by_brand_title)

	v1.POST("/catalog-items", items.Create_item)
	v1.PUT("/catalog-items", items.Update_item)
	v1.DELETE("/catalog-items/:id", items.Delete_item)

	v2 := r.Group("/api/v2")

	v2.GET("/catalog-items", items_v2.Catalog_items)

}
