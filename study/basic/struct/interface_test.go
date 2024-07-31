package _struct

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"testing"
)

type testInterface1 interface {
	M1()
	M2()
}

type testStruct1 struct {
	testInterface1
}

func Test()  {

}

func TestMethodSet(t *testing.T) {
	var struct1 = &testStruct1{}
	struct1.M1()
	struct1.M2()
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	r.width = 10
	return r.width * r.height
}
func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	c.radius = 10
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterface(t *testing.T) {
	var r geometry = &rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	fmt.Println(r)
	measure(c)
	fmt.Println(c)
}

func TestAssert(t *testing.T) {
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64)
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type o
	v3 := i.(int64)
	fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is i
	v4 := i.([]int)                                     // panic: interface conversion: interface {} is int64, not []int
	fmt.Printf("the type of v4 is %T\n", v4)
}

func TestTypeSwitch(t *testing.T) {
	var x interface{} = 13
	switch v := x.(type) {
	case nil:
		println("v is nil")
	case int:
		println("the type of v is int, v =", v)
	case string:
		println("the type of v is string, v =", v)
	case bool:
		println("the type of v is bool, v =", v)
	default:
		println("don't support the type")
	}
}

type MyError struct {
	message string
}

func (m *MyError) Error() string {
	return m.message
}

func TestError1(t *testing.T) {
	error1 := errors.New("123")
	error2 := fmt.Errorf("fmt create error")
	error3 := fmt.Errorf("fmt create error2 ,err:%w", error2)
	fmt.Println(error1)
	fmt.Println(error2)
	fmt.Println(error3)
	//实现自定义error 狠简单就是实现error接口就可以
	var myError error = &MyError{
		message: "1231231",
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
		message: "1231231",
	}
	var err error = fmt.Errorf("%w", myError)
	//func Unwrap(err error) error                 // 获得err包含下一层错误
	fmt.Println(errors.Unwrap(myError))
	//func Is(err, target error) bool              // 判断err是否包含target
	fmt.Println(errors.Is(myError, err))
	//func As(err error, target interface{}) bool  // 判断err是否为target类型
	//fmt.Println(errors.As(myError, err))
}
