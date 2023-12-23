package book

import (
	"demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBook(c *gin.Context) {
	c.JSON(http.StatusOK, common.Ok("bookInfo"))
}

func getBooks(c *gin.Context) {

}
