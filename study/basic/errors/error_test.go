package errors

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Printf("%+v", throw())
}

func throw() error {
	return errors.New("error")
}

func Test1(t *testing.T) {
	//if err != nil {
	//	// 使用 %w 指令，会返回一个包装了err的错误，接收方可以从父error中获取到源error
	//	// 并根据判断错误是否为某一个错误类型
	//	// sourceErr --wrap--> WrapErr(sourceErr)
	//	return fmt.Errorf("bar failed: %w", err)
	//}
	//
	//if err != nil {
	//	// 使用 %v 指令，不会包装错误，会直接转化为另一个错误，源错误不再可用
	//	// sourceErr --transform--> otherErr
	//	return fmt.Errorf("bar failed: %v", err)
	//}
}

// 一种解决方式是使用命名结果，把错误信息通过命名结果返回
func getBalance(db *sql.DB, clientID string) (balance float32, err error) {
	query := "select * from table1"
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil {
			if closeErr != nil {
				// 如果db.Query语句也出现了执行错误，把close错误的信息打印出来
				log.Printf("failed to close rows: %v", err)
			}
			return
		}
		// 当只要close的时候有错误，把closeErr赋值给err传递到上一层
		err = closeErr
	}()

	// Use rows
	return 0, nil
}
