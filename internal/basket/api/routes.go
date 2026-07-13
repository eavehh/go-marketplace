package api

import (
	"github.com/eavehh/marketpl.microserv/internal/basket/api/handlers"
	"github.com/gin-gonic/gin"
)

func Register_routes(r *gin.Engine,
	cart_handler *handlers.Cart_handler,
) {
	v1 := r.Group("/api/v1")

	v1.POST("/cart", cart_handler.Save_cart)
	v1.GET("/cart/:account_name", cart_handler.Get_cart)
}
