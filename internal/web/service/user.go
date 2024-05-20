package service

import (
	model2 "demo/internal/web/model"
	store2 "demo/internal/web/store"
)

type UserSrv interface {
	Get(id int64) *model2.UserInfo
}

var _ UserSrv = (*userService)(nil)

type userService struct {
	uDao store2.UserRepository
}

func (u *userService) Get(id int64) *model2.UserInfo {
	return u.uDao.GetById(id).Info()
}
