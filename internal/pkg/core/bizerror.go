package core

import (
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

type BizError struct {
	Code       int
	Message    string
	HttpStatus int
}

func (b *BizError) Error() string {
	return b.Message
}
func ErrorFmt(err error, printError bool) *BizError {
	if err == nil {
		return Success
	}
	if printError {
		log.Error(err)
	}
	switch typed := err.(type) {
	case *BizError:
		return typed
	default:
		return &BizError{Code: http.StatusInternalServerError, Message: "服务器内部错请稍后再试", HttpStatus: http.StatusInternalServerError}
	}
}
