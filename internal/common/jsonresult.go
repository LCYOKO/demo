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
		Timestamp: time.Now().UnixNano(),
	}
}

func Ok(data interface{}) *JsonResult {
	return &JsonResult{
		Code:      OK.GetCode(),
		Msg:       OK.GetMsg(),
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}
}

func DefaultError() *JsonResult {
	return &JsonResult{
		Code:      ERROR.GetCode(),
		Msg:       ERROR.GetMsg(),
		Timestamp: time.Now().UnixNano(),
	}
}

func Error(code BizCode, msg string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:      code.GetCode(),
		Msg:       msg,
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}
}
