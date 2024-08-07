package goin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestRedirect(t *testing.T) {
	engine.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	engine.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "test1"
		engine.HandleContext(c)
	})
	engine.Run(":8080")
}

func TestCros(t *testing.T) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},                         // 允许跨域发来请求的网站
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
