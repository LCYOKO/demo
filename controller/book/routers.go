package book

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	group := e.Group("/books")
	{
		group.GET("/info", getBook)
		group.GET("/infos", getBooks)
	}
}
