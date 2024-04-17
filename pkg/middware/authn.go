// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package middleware

import (
	"demo/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/miniblog/pkg/token"
	"go.uber.org/zap"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// 如果合法则将 token 中的 sub 作为<用户名>存放在 gin.Context 的 XUsernameKey 键中.
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(c)
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
