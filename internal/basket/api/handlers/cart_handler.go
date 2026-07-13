package handlers

import (
	"fmt"
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/basket/application/commands"
	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
	"github.com/gin-gonic/gin"
)

type Cart_handler struct {
	save_cart *commands.Save_cart_handler
}

func New_cart_handler(save_cart *commands.Save_cart_handler) *Cart_handler {
	return &Cart_handler{save_cart: save_cart}
}

func (h *Cart_handler) Save_cart(c *gin.Context) {
	var req struct {
		Cart domain.Shopping_cart `json:"cart"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}

	account_name, err := h.save_cart.Handle(
		c.Request.Context(),
		commands.Save_cart_command{Cart: req.Cart},
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_name": account_name,
		"location":     fmt.Sprintf("%s/%s", c.FullPath(), account_name),
	})
}
