package web

//var (
//	// ErrSignToken 表示签发 JWT Token 时出错.
//	ErrSignToken = &Errno{HTTP: 401, Code: "AuthFailure.SignTokenError", Message: "Error occurred while signing the JSON web token."}
//
//	// ErrTokenInvalid 表示 JWT Token 格式错误.
//	ErrTokenInvalid = &Errno{HTTP: 401, Code: "AuthFailure.TokenInvalid", Message: "Token was invalid."}
//
//	// ErrUnauthorized 表示请求没有被授权.
//	ErrUnauthorized = &Errno{HTTP: 401, Code: "AuthFailure.Unauthorized", Message: "Unauthorized."}
//)

var (
	// OK 代表请求成功.
	OK = &BizCode{
		200,
		200,
		"请求成功",
	}
	// InvalidParam 表示参数绑定错误.
	InvalidParam = &BizCode{
		400,
		400,
		"参数校验异常",
	}
	Unauthorized = &BizCode{
		401,
		401,
		"身份未认证",
	}
	// TokenInvalid JWT Token 格式错误
	TokenInvalid = &BizCode{
		401,
		401,
		"token非法",
	}
	// NotPermission 没有权限
	NotPermission = &BizCode{
		403,
		403,
		"没有权限",
	}
	// PageNotFound 表示路由不匹配错误.
	PageNotFound = &BizCode{
		404,
		404,
		"请求资源不存在",
	}
	// InternalServerError 表示所有未知的服务器端错误.
	InternalServerError = &BizCode{
		500,
		500,
		"服务器内部错误",
	}
)

type BizCode struct {
	Status int
	Code   int
	Msg    string
}

//// Decode 尝试从 err 中解析出业务错误码和错误信息.
//func Decode(err error) (int, int, string) {
//	if err == nil {
//		return OK.Status, OK.Code, OK.Msg
//	}
//
//	switch typed := err.(type) {
//	case *BizError:
//		return typed.Status, typed.Code, typed.Msg
//	default:
//	}
//	// 默认返回未知错误码和错误信息. 该错误代表服务端出错
//	return InternalServerError.Status, InternalServerError.Code, InternalServerError.Msg
//}
