package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (u *Controller) GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "UserInfo")
}

func (u *Controller) GetInfos(c *gin.Context) {
	//userInfos := make([]model.UserInfo,0)
	//dao.DB_Test.Table("user_info").Find(&userInfos)
	//c.JSON(http.StatusOK,userInfos)
}
