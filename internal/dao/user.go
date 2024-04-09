package dao

type UserDao interface {
	GetById(id int64) *User
	GetByIds(ids []int64) []*User
}
