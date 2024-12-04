package middleware

import (
	"demo/pkg/auth"
	"demo/pkg/web"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// 如果合法则将 token 中的 sub 作为<用户名>存放在 gin.Context 的 XUsernameKey 键中.
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := auth.ParseRequest(c)
		if err != nil {
			zap.L().Error("Parse jtwToken failed.", zap.Error(err))
			web.WriteResponse(c, web.TokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set(web.XUsernameKey, username)
		c.Next()
	}
}
