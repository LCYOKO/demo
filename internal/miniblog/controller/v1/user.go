package v1

import (
	"demo/internal/miniblog/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc service.UserService
}

func (u *UserController) GetById(ctx *gin.Context) {
	//u.userSvc.GetById()
}
