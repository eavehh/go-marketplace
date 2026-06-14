package handlers

import (
	"net/http"

	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type Brands_handler struct {
	brands *queries.Brands_handler
}

func New_brands_handler(brands *queries.Brands_handler) *Brands_handler {
	return &Brands_handler{brands: brands}
}

func (h *Brands_handler) Brands(c *gin.Context) {
	brands, err := h.brands.Handle(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"brands": brands,
	})
}
