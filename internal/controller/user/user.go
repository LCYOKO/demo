package user

import (
	"demo/internal/service"
	"demo/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	usrv service.UserSrv
}

func (u *Controller) GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, web.Ok(nil))
}

func (u *Controller) GetInfos(c *gin.Context) {
	//userInfos := make([]model.UserInfo,0)
	//store.DB_Test.Table("user_info").Find(&userInfos)
	//c.JSON(http.StatusOK,userInfos)
}
