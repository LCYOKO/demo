package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getBook(c *gin.Context) {
	time.Sleep(time.Second * 10)
	c.JSON(http.StatusOK, "BookInfo")
}

func getBooks(c *gin.Context) {

}
