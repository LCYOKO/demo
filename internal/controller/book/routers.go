package book

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var bookController *Controller

func Instance() *Controller {
	once.Do(func() {
		bookController = &Controller{}
	})
	return bookController
}

func Routers(e *gin.Engine) {
	group := e.Group("/books")
	{
		bookController = Instance()
		group.GET("/info", bookController.GetBook)
		group.GET("/infos", bookController.GetBooks)

	}
}
