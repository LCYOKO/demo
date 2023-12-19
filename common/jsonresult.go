package common

import "time"

type JsonResult struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func DefaultOk() *JsonResult {
	return &JsonResult{
		Code:      OK.GetCode(),
		Msg:       OK.GetMsg(),
		Timestamp: time.Now().UnixMilli(),
	}
}

func Ok(data interface{}) *JsonResult {
	return &JsonResult{
		Code:      OK.GetCode(),
		Msg:       OK.Msg,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
	}
}

func DefaultError() *JsonResult {
	return &JsonResult{
		Code:      ERROR.Code,
		Msg:       ERROR.Msg,
		Timestamp: time.Now().UnixMilli(),
	}
}

func Error(code BizCode, msg string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:      code.GetCode(),
		Msg:       msg,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
	}
}
