package controller

import (
	book2 "demo/internal/web/controller/book"
	user2 "demo/internal/web/controller/user"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init(g *gin.Engine) error {
	book2.Routers(g)
	user2.Routers(g)
	return nil
}
