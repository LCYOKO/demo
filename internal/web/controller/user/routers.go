package user

import (
	"github.com/gin-gonic/gin"
)

var userController *Controller

func Routers(g *gin.Engine) {
	userController = &Controller{}
	group := g.Group("/user")
	{
		group.GET("/info", userController.GetInfo)
		group.GET("/infos", userController.GetInfos)
	}
}
