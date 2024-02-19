package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getInfo(c *gin.Context)  {
	c.JSON(http.StatusOK,"UserInfo")
}

func getInfos(c *gin.Context)  {
	//userInfos := make([]model.UserInfo,0)
	//dao.DB_Test.Table("user_info").Find(&userInfos)
	//c.JSON(http.StatusOK,userInfos)
}
