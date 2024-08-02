package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,)
	})
	engine.Run(":8080")
}
