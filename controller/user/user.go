package user

import (
	"demo/dao"
	"demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(e *gin.Engine){
	group := e.Group("/user")
	{
		group.GET("/info",getInfo)
        group.GET("/infos",getInfos)
	}
}

func getInfo(c *gin.Context)  {
	c.JSON(http.StatusOK,"UserInfo")
}

func getInfos(c *gin.Context)  {
	userInfos := make([]model.UserInfo,0)
	dao.DB_Test.Table("user_info").Find(&userInfos)
	c.JSON(http.StatusOK,userInfos)
}
