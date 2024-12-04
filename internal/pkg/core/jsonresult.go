package core

import (
	"github.com/gin-gonic/gin"
	"time"
)

type JsonResult struct {
	Code      int
	Message   string
	data      interface{}
	timestamp int64
}

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	code := ErrorFmt(err, true)
	ctx.JSON(code.HttpStatus, JsonResult{
		Code:      code.Code,
		Message:   code.Message,
		data:      data,
		timestamp: time.Now().UnixNano(),
	})
}
