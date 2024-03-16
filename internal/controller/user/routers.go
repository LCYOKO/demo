package user

import "github.com/gin-gonic/gin"

func Route(e *gin.Engine) {
	group := e.Group("/user")
	{
		group.GET("/info", getInfo)
		group.GET("/infos", getInfos)
	}
}
