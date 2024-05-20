package user

import (
	service2 "demo/internal/web/service"
	"demo/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	usrv service2.UserSrv
}

func (u *Controller) GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, web.Ok(nil))
}

func (u *Controller) GetInfos(c *gin.Context) {
	//userInfos := make([]model.UserInfo,0)
	//store.DB_Test.Table("user_info").Find(&userInfos)
	//c.JSON(http.StatusOK,userInfos)
}
