package dao

import "gorm.io/gorm"

type UserDao interface {
	GetById(id int64) *User
	GetByIds(ids []int64) []*User
}

type userDaoImpl struct {
	db *gorm.DB
}

func (u *userDaoImpl) GetById(id int64) *User {
	return nil
}

func (u *userDaoImpl) GetByIds(ids []int64) []*User {
	return nil
}
