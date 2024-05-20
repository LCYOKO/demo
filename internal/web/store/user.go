package store

import (
	model2 "demo/internal/web/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id int64) *model2.User
	GetByIds(ids []int64) []*model2.User
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) GetById(id int64) *model2.User {
	return nil
}

func (u *userRepository) GetByIds(ids []int64) []*model2.User {
	return nil
}
