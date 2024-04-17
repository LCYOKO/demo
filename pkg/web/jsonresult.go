package web

import (
	"github.com/gin-gonic/gin"
	"time"
)

type JsonResult struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func DefaultOk() *JsonResult {
	return &JsonResult{
		Code:      OK.Code,
		Msg:       OK.Msg,
		Timestamp: time.Now().UnixNano(),
	}
}

func Ok(data interface{}) *JsonResult {
	return &JsonResult{
		Code:      OK.Code,
		Msg:       OK.Msg,
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}
}

func DefaultError() *JsonResult {
	return &JsonResult{
		Code:      InternalServerError.Code,
		Msg:       InternalServerError.Msg,
		Timestamp: time.Now().UnixNano(),
	}
}

func Error(code int, msg string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:      code,
		Msg:       msg,
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}
}

func Of(bizCode *BizCode, data interface{}) *JsonResult {
	return &JsonResult{
		Code: bizCode.Code,
		Msg:  bizCode.Msg,
		Data: data,
	}
}

func WriteResponse(c *gin.Context, code *BizCode, data interface{}) {
	c.JSON(code.Status, Of(code, data))
}
