package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/smartystreets/goconvey/convey"
	//"bou.ke/monkey"
)

var DB_Test *gorm.DB

func Init()  error{
	dsn := "root:3852159@tcp(localhost:3306)/miaosha?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		return err
	}
	DB_Test =db
	return nil
}
