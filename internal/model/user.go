package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

type UserInfo struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Age          int32
	TelPhone     string
	RegisterMode string
	ThirdPartId  int64
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Info() *UserInfo {
	return nil
}
