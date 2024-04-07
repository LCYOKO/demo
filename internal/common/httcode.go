package common

var (
	OK = &BizCode{
		200,
		"请求成功",
	}
	NOT_FOUND = &BizCode{
		404,
		"请求资源不存在",
	}
	ERROR = &BizCode{
		500,
		"服务器内部错误",
	}
)

type BizCode struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

func (b *BizCode) GetCode() int {
	return b.code
}

func (b *BizCode) GetMsg() string {
	return b.msg
}
