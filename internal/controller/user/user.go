package user

import (
	"demo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	usrv service.UserSrv
}

func (u *Controller) GetInfo(c *gin.Context) {
	 c.Query("id")
	c.JSON(http.StatusOK, u.usrv.Get(1))
}

func (u *Controller) GetInfos(c *gin.Context) {
	//userInfos := make([]model.UserInfo,0)
	//dao.DB_Test.Table("user_info").Find(&userInfos)
	//c.JSON(http.StatusOK,userInfos)
}
