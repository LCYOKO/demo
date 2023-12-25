package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

func (u *User) TableName() string {
	return "user"
}

// https://gorm.io/docs/index.html
func init()  {
	var err error
	dsn := "root:lcyoko@tcp(114.55.147.178:33060)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect mysql failed")
	}
}

func TestCreate() {
	u := &User{
		Name: "lisi",
	}
	create := db.Create(u)
	fmt.Println("u", create.RowsAffected)
	var user User
	db.First(&user)
	fmt.Println("user", user)
}
