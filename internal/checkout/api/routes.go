package api

import (
	"github.com/eavehh/marketpl.microserv/internal/checkout/api/handlers"
	"github.com/gin-gonic/gin"
)

func Register_routes(r *gin.Engine, order_handler *handlers.Order_handler) {
	api := r.Group("/api/v1")
	{
		orders := api.Group("/orders")
		{
			orders.GET("/:id", order_handler.Order)
		}
	}
}
