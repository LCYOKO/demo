package book

import (
	common2 "demo/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (b *Controller) GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, common2.Ok("bookInfo"))
}

func (b *Controller) GetBooks(c *gin.Context) {

}
