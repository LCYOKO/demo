package book

import (
	"demo/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (b *Controller) GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, common.Ok("bookInfo"))
}

func (b *Controller) GetBooks(c *gin.Context) {

}
