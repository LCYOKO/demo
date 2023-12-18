package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Route(e *gin.Engine)  {
	group := e.Group("/book")
	{
		group.GET("/info",getBook)
		group.GET("/infos",getBooks)
	}
}

func getBook(c *gin.Context){
	time.Sleep(time.Second*10)
   c.JSON(http.StatusOK,"BookInfo")
}

func getBooks(c *gin.Context)  {

}



