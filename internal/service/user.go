package service

import (
	"demo/internal/dao"
	"demo/internal/model"
)

type UserSrv interface {
	Get(id int64) *model.UserInfo
}

type userService struct {
	uDao dao.UserRepository
}

func (u *userService) Get(id int64) *model.UserInfo {
	return u.uDao.GetById(id).Info()
}
