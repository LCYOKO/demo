package model

import (
	"database/sql"
)

type User struct {
	Id        sql.NullInt64 `gorm:"primaryKey"`
	Name      sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Info() *UserInfo {
	return &UserInfo{Id: u.Id.Int64, Name: u.Name.String}
}

type UserInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
