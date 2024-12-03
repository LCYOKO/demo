package errors

// 你应该只处理一次错误。处理一个错误意味着检查错误值，并做出单一的决定。
//
import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestError1(t *testing.T) {
	error1 := errors.New("123")
	error2 := fmt.Errorf("fmt create error")
	error3 := fmt.Errorf("fmt create error2 ,err:%w", error2)
	fmt.Println(error1)
	fmt.Println(error2)
	fmt.Println(error3)
	//实现自定义error 狠简单就是实现error接口就可以
	var myError error = &MyError{
		err: "1231231",
	}
	fmt.Println(myError)
}

func TestError2(t *testing.T) {
	// 使用pkg.errors打印当前的协程栈
	fmt.Printf("%+v", createError())
}

func createError() error {
	return errors.New("1231")
}

func TestError3(t *testing.T) {
	var myError error = &MyError{
		err: "1231231",
	}
	var err error = fmt.Errorf("%w", myError)
	//func Unwrap(err error) error                 // 获得err包含下一层错误
	fmt.Println(errors.Unwrap(myError))
	//func Is(err, target error) bool              // 判断err是否包含target
	fmt.Println(errors.Is(myError, err))
	//func As(err error, target interface{}) bool  // 判断err是否为target类型
	//fmt.Println(errors.As(myError, err))
}

func TestError(t *testing.T) {
	fmt.Printf("%+v", throw())
}

func throw() error {
	return errors.New("error")
}

func test(err error) error {
	if err != nil {
		// 使用 %w 指令，会返回一个包装了err的错误，接收方可以从父error中获取到源error
		// 并根据判断错误是否为某一个错误类型
		// sourceErr --wrap--> WrapErr(sourceErr)
		return fmt.Errorf("bar failed: %w", err)
	}

	if err != nil {
		// 使用 %v 指令，不会包装错误，会直接转化为另一个错误，源错误不再可用
		// sourceErr --transform--> otherErr
		return fmt.Errorf("bar failed: %v", err)
	}
	return nil
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

type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return e.err
}

func TestError4(t *testing.T) {
	if errors.New("error") == errors.New("error") {
		fmt.Println("equals1")
	}
	e1 := &MyError{err: "err"}
	e2 := &MyError{err: "err"}
	if e1 == e2 {
		fmt.Println("equals2")
	}
}

func TestError5(t *testing.T) {
	var err error = &MyError{}
	println(err.Error())
	ErrorsPkg()
}

func ErrorsPkg() {
	err := &MyError{}
	// 使用 %w 占位符，返回的是一个新错误
	// wrappedErr 是一个新类型，fmt.wrapError
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	// 再解出来
	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped")
	}

	if errors.Is(wrappedErr, err) {
		// 虽然被包了一下，但是 Is 会逐层解除包装，判断是不是该错误
		fmt.Println("wrapped is err")
	}
	copyErr := &MyError{}
	// 这里尝试将 wrappedErr转换为 MyError
	// 注意我们使用了两次的取地址符号
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error")
	}
}

// ### 重复 `Wrap()` 的坑
// 使用 `Wrap()`虽然打印日志时很方便的附带堆栈信息，但使用时也有一个不小的坑：
// - 就是**多处 `Wrap()`**，导致打印的错误时会有**多倍的堆栈信息**，因为每 `Wrap()`一次，底层便会调用一次 `withStack()`，就会多输出一次堆栈信息。
//
// 如何合理的使用 `Wrap()`呢，给出以下几点建议：
// - 1. 在你的**应用代码(指偏向业务的逻辑，不是封装的基础库)** 中，使用 `errors.New` 或者  `errros.Errorf` 返回自定义错误，注意都是指 `pkg/errors`库，如：
// ```go
//
//	func parseArgs(args []string) error {
//		if len(args) < 3 {
//			return erros.Errorf("not enough arguments, expected at least 3 argument")
//		}
//	}
//
// ```
// - 2. 如果**调用其他包内的函数（即项目某个函数）**，通常简单的**直接返回err**，如：
//
//	if err := somePkg.Logic();err != nil{
//		return err
//	}
//
// - 3. 如果是**最底层的业务逻辑，通常是与数据库相关的**，考虑使用 `errors.Wrap` 或者 `errors.Wrapf` **包装**数据库返回的err
// - 4. 如果**和第三方库(如github这类库)、标准库、公司或个人封装的基础库进行协作**，考虑使用 `errors.Wrap` 或者 `errors.Wrapf` **包装**这些库返回的err，如：
// f, err := os.Open(filePath)
//
//	if err != nil {
//		return errros.Wrapf(err, "failed to open %q", filePath)
//	}
//
// 记不住上面具体建议没关系，记住一个基本原则：**最底层的逻辑需要wrap**，如业务开发是数据库操作相关（Mysql、MongoDB、Redis等），调用基础库，如GO标准库或第三方库
func TestWrapError(t *testing.T) {
	_, err := ReadFile("test")
	fmt.Printf("err:%+v", err)
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		// Wrapf err： 包含堆栈信息，可以格式化错误内容
		return nil, errors.Wrapf(err, "failed to open %q", path) // %q 单引号围绕的字符字面值，由Go语法安全地转义，这样中文文件名也能正确显示
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		// Wraperr： 包含堆栈信息
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func TestErrWithMessage(t *testing.T) {
	_, err := ReadConfig("test")
	fmt.Printf("err:%+v", err)
}

var filePath = "test_file_path"

func ReadConfig(path string) ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".yaml"))
	return config, errors.WithMessage(err, "could mot read config")
}
