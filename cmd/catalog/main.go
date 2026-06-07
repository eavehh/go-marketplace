package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	if err := engine.Run(":9001"); err != nil {
		log.Fatalf(err.Error())
	}
}
