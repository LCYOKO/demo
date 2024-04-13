package dao

import (
	"demo/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id int64) *model.User
	GetByIds(ids []int64) []*model.User
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) GetById(id int64) *model.User {
	return nil
}

func (u *userRepository) GetByIds(ids []int64) []*model.User {
	return nil
}
