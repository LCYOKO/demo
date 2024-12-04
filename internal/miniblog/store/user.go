package store

import (
	"demo/internal/miniblog/model"
	"demo/internal/miniblog/store/mysql"
)

type UserRepository interface {
	GetById(id int64) *model.User
	GetByIds(ids []int64) []*model.User
	Update(user *model.User) int
	deleteById(id int64) int
}

type userRepositoryImpl struct {
	userStore *store.Datastore
}

var _ UserRepository = (*userRepositoryImpl)(nil)

func (u *userRepositoryImpl) GetById(id int64) *model.User {
	return nil
}

func (u *userRepositoryImpl) GetByIds(ids []int64) []*model.User {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryImpl) Update(user *model.User) int {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryImpl) deleteById(id int64) int {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		userStore: store.Store,
	}
}
