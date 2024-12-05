package service

import (
	"demo/internal/miniblog/model"
	"demo/internal/miniblog/store"
)

type UserService interface {
	Login() error
	GetById(id int64) (*model.UserInfo, error)
	GetByIds(ids []int64) ([]model.UserInfo, error)
	Update() error
	DeleteById(id int64) error
}

var _ UserService = (*userServiceImpl)(nil)

type userServiceImpl struct {
	userRpo store.UserRepository
}

func (u userServiceImpl) Login() error {
	//TODO implement me
	panic("implement me")
}

func (u userServiceImpl) GetById(id int64) (*model.UserInfo, error) {
	user, err := u.userRpo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user.Info(), nil
}

func (u userServiceImpl) GetByIds(ids []int64) ([]model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceImpl) Update() error {
	return nil
}

func (u userServiceImpl) DeleteById(id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService() UserService {
	return &userServiceImpl{
		userRpo: store.NewUserRepository(),
	}
}
