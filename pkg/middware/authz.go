package middleware

import (
	"demo/pkg/web"
	"github.com/gin-gonic/gin"
)

// Auther 用来定义授权接口实现.
// sub: 操作主题，obj：操作对象, act：操作
type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是 Gin 中间件，用来进行请求授权.
func Authz(a Auther) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString(web.XUsernameKey)
		obj := c.Request.URL.Path
		act := c.Request.Method

		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			web.WriteResponse(c, web.Unauthorized, nil)
			c.Abort()
			return
		}
	}
}
