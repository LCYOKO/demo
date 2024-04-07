package user

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var userController *Controller

func Route(e *gin.Engine) {
	userController = Instance()
	group := e.Group("/user")
	{
		group.GET("/info", userController.GetInfo)
		group.GET("/infos", userController.GetInfos)
	}
}

func Instance() *Controller {
	once.Do(func() {
		userController = &Controller{}
	})
	return userController
}
