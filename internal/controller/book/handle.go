package book

import (
	common2 "demo/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBook(c *gin.Context) {
	c.JSON(http.StatusOK, common2.Ok("bookInfo"))
}

func getBooks(c *gin.Context) {

}
