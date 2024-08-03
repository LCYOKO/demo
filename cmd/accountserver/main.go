package main

import (
	"demo/internal/accountserver/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required" message:"password should not empty" `
}

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, &model.Account{Id: 1, Name: "lisi"})
	})
	engine.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if err := c.ShouldBind(&form); err == nil {
			fmt.Println(form)
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
		c.Err()
	})
	engine.Run(":8080")
}
