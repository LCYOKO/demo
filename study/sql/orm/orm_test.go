package orm

// https://gorm.io/zh_CN/docs/query.html
// https://liwenzhou.com/posts/Go/gorm-crud/

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"testing"
	"time"
)

var db *gorm.DB

type User struct {
	Id   int64
	Name sql.NullString
	Age  int64
	//CreateTime time.Time
	//UpdateTime sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

func setUp() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:lcyoko123@tcp(127.0.0.1:33060)/test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                              // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                             // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                             // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                             // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                            // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("error", err)
	}
}

func TestMain(t *testing.M) {
	setUp()
	code := t.Run()
	os.Exit(code)
}

func TestInsert(t *testing.T) {
	u := &User{Age: 1}
	db.Create(u)
	fmt.Printf("u %v\n", u)
}

func TestQuery(t *testing.T) {

	user := &User{}
	db.Find(user, 1)
	fmt.Println(user)

	users := make([]User, 3)
	db.Find(&users, []int{1, 5, 6})
	fmt.Println(users)

	db.Where("name is null").Find(&users)
	fmt.Println(users)

	// 为查询 SQL 添加额外的 SQL 操作
	db.Set("gorm:query_option", "FOR UPDATE").Find(&user, 10)

	db.Select("age").Where("name is null").Limit(1).Find(&users)
	fmt.Println(users)
}

func TestUpdate(t *testing.T) {
	db.Model(&User{}).Where("id = ?", 6).Update("name", "hello")

	users := make([]User, 3)
	db.Find(&users, []int{1, 5, 6})
	fmt.Println(users)

	db.Where("name is null").Find(&users)
	fmt.Println(users)
}

func TestDelete(t *testing.T) {
	db.Model(&User{}).Delete("id = ?", 6)

	users := make([]User, 3)
	db.Find(&users, []int{1, 5, 6})
	fmt.Println(users)
}
