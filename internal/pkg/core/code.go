package core

var (
	Success = &BizError{Code: 0, Message: "success", HttpStatus: 200}

	ParamError = &BizError{Code: 2, Message: "参数错误", HttpStatus: 400}

	AuthError = &BizError{Code: 3, Message: "未授权", HttpStatus: 401}

	Forbidden = &BizError{Code: 4, Message: "禁止访问", HttpStatus: 403}

	NotFound = &BizError{Code: 5, Message: "找不到资源", HttpStatus: 404}

	ServerError = &BizError{Code: 6, Message: "服务器错误", HttpStatus: 500}
)
