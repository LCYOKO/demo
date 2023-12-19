package common

var (
	// HTTP原生状态码
	OK = &BizCode{
		200,
		"请求成功",
	}
	NOT_FOUND = &BizCode{
		404,
		"请求也",
	}
	ERROR = &BizCode{
		500,
		"服务器内部错误",
	}
)

type BizCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (b BizCode) GetCode() int {
	return b.Code
}

func (b BizCode) GetMsg() string {
	return b.Msg
}
