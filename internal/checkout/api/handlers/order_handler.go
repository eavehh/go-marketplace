package handlers

import (
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/checkout/application/queries"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Order_handler struct {
	order_by_id_query_handler queries.Order_by_id_query_handler
}

func New_order_handle(order_by_id *queries.Order_by_id_query_handler) *Order_handler {
	return &Order_handler{
		order_by_id_query_handler: *order_by_id,
	}
}

func (h *Order_handler) Order(c *gin.Context) {
	id_str := c.Param("id")
	id, err := uuid.Parse(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	query := queries.Order_by_id_query{Id: id}
	order, err := h.order_by_id_query_handler.Handle(c, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{

			"error": err.Error(),
		})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "order not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": order,
	})

}
