package testgorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Id   int64
	Name string
	Age  int
}

// https://gorm.io/docs/index.html
func init() {
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
	db.Create(u)

	var user User
	db.First(&user)
	fmt.Println(user)
}
