package api

import (
	"github.com/eavehh/marketpl.microserv/internal/catalog/api/handlers"
	"github.com/gin-gonic/gin"
)

func Register_routes(r *gin.Engine,
	brands *handlers.Brands_handler,
	categories *handlers.Categories_handler,
) {
	v1 := r.Group("/api/v1")

	v1.GET("/brands", brands.Brands)
	v1.GET("categories", categories.Categories)
}
