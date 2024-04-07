package service

import "demo/internal/model"

type UserSrv interface {
	GetUser() *model.UserInfo
}

type UseService struct {
}

func (u *UseService) GetUser() *model.UserInfo {
	return nil
}
