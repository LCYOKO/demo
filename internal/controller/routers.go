package controller

import (
	"demo/internal/controller/book"
	"demo/internal/controller/user"
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
	book.Routers(g)
	user.Routers(g)
	return nil
}
