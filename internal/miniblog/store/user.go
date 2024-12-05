package store

import (
	"demo/internal/miniblog/model"
	"demo/internal/miniblog/store/mysql"
	"errors"
)

type UserRepository interface {
	GetById(id int64) (*model.User, error)
	GetByIds(ids []int64) ([]model.User, error)
	Update(user *model.User) (int64, error)
	deleteById(id int64) (int64, error)
}

type userRepositoryImpl struct {
	store *store.Datastore
}

var _ UserRepository = (*userRepositoryImpl)(nil)

func (u *userRepositoryImpl) GetById(id int64) (*model.User, error) {
	var user model.User
	find := u.store.UserDb.Find(&user, id)
	if find.Error != nil {
		return nil, find.Error
	}
	return &user, nil
}

func (u *userRepositoryImpl) GetByIds(ids []int64) ([]model.User, error) {
	if len(ids) == 0 {
		return []model.User{}, nil
	}
	users := make([]model.User, len(ids))
	find := u.store.UserDb.Find(&users, ids)
	if find.Error != nil {
		return nil, find.Error
	}
	return users, nil
}

func (u *userRepositoryImpl) Update(user *model.User) (int64, error) {
	if user == nil {
		return 0, nil
	}
	if !user.Id.Valid {
		return 0, errors.New("id not set")
	}
	result := u.store.UserDb.Save(user)
	return result.RowsAffected, result.Error
}

func (u *userRepositoryImpl) deleteById(id int64) (int64, error) {
	result := u.store.UserDb.Delete(&model.User{}, id)
	return result.RowsAffected, result.Error
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		store: store.Store,
	}
}
