package goin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestMiddleWare(t *testing.T) {
	StatCost := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			start := time.Now()
			// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
			c.Set("name", "小王子")
			// 调用该请求的剩余处理程序
			c.Next()
			// 不调用该请求的剩余处理程序
			// c.Abort()
			// 计算耗时
			cost := time.Since(start)
			fmt.Println(cost)
		}
	}
	// 注册一个全局中间件
	engine.Use(StatCost())
	engine.GET("/cost", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		fmt.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	engine.Run(":8080")
}
