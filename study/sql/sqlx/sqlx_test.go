package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

var db *sqlx.DB

func initDB() {
	var err error
	dsn := "user:password@tcp(114.55.147.178:33060)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		panic("connect sql error")
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}

func TestQuery(t *testing.T) {
	initDB()
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("user%#v\n", u)
	//--------------------------
	sqlStr = "select id, name, age from user where id > ?"
	var users []User
	err = db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func TestInsert(t *testing.T) {
	initDB()
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func TestDelete(t *testing.T) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func TestUpdate(t *testing.T) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func TestBatchInsert(t *testing.T) {
	initDB()

	users := []interface{}{
		User{Name: "七米", Age: 18},
		User{Name: "q1mi", Age: 28},
		User{Name: "小王子", Age: 38},
	}
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query)
	fmt.Println(args)
	db.Exec(query, args...)
}

func TestBatchInsert2(t *testing.T) {
	initDB()
	users := []User{
		{

		},
		{

		},
	}
	_, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	if err != nil {
		t.Log("")
	}
}

func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func TestTransactionDemo2(t *testing.T) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			// re-throw panic after Rollback
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			// err is non-nil; don't change it
			_ = tx.Rollback()
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return
	}
	if n != 1 {
		return
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return
	}
	if n != 1 {
		return
	}
}

type User struct {
	Id   int64
	Name string
	Age  int
}
