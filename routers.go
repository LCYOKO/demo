package main

import (
	"demo/dao"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	dao.Init()
	r := gin.Default()
	for _, opt := range options {
		opt(r)
	}
	return r
}
