package service

import (
	"demo/internal/dao"
	"demo/internal/model"
)

type UserSrv interface {
	Get(id int64) *model.UserInfo
}

type UseService struct {
	uDao dao.UserDao
}

func (u *UseService) Get(id int64) *model.UserInfo {
	return u.uDao.GetById(id)
}
