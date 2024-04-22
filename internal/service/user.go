package service

import (
	"demo/internal/model"
	"demo/internal/store"
)

type UserSrv interface {
	Get(id int64) *model.UserInfo
}

var _ UserSrv = (*userService)(nil)

type userService struct {
	uDao store.UserRepository
}

func (u *userService) Get(id int64) *model.UserInfo {
	return u.uDao.GetById(id).Info()
}
