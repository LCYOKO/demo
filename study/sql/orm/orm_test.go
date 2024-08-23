package orm

// https://gorm.io/zh_CN/docs/query.html
// https://liwenzhou.com/posts/Go/gorm-crud/

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var db *gorm.DB

type User struct {
	Id         int64
	Name       string
	Age        int64
	CreateTime time.Time
	UpdateTime sql.NullTime
}

func (u User) TableName() string {
	return "user"
}

func setUp() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:lcyoko123@tcp(127.0.0.1:33060)/users?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                               // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
	}
}

func TestMain(t *testing.M) {
	setUp()
	code := t.Run()
	os.Exit(code)
}

func TestQuery(t *testing.T) {
	user := &User{}
	db.Find(user, 1)
	fmt.Println(user)
	users := make([]User, 3)
	result := db.Find(users, []int{1, 2, 3})
	fmt.Println(users, result)
}

//
//func TestInsert(t *testing.T) {
//	user := &model2.User{
//		Name: "wangwu",
//	}
//	result := db.Create(user)
//	if result.Error != nil {
//		fmt.Printf("insert error %v\n", result.Error)
//	} else {
//		fmt.Printf("insert success %v\n", result.RowsAffected)
//	}
//}
