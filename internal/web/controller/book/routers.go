package book

import (
	"github.com/gin-gonic/gin"
)

var bookController *Controller

func Routers(g *gin.Engine) {
	group := g.Group("/books")
	{
		bookController = &Controller{}
		group.GET("/info", bookController.GetBook)
		group.GET("/infos", bookController.GetBooks)

	}
}
