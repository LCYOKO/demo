package service

import (
	"demo/internal/web/store"
	v1 "demo/pb/miniblog/v1"
)

type UserService interface {
	Login() (result bool, err error)
	GetById(id int64) *v1.UserInfo
	GetByIds(ids []int64) []*v1.UserInfo
	Update()
	DeleteById(id int64) bool
}
type userServiceImpl struct {
	userRpo store.UserRepository
}
