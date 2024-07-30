package basic

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	//定义数组
	var arr1 [4]int
	var arr2 = [4]int{1, 2, 3, 4}
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr1, arr2, arr3)
	modifyArray(&arr3)
	fmt.Println(arr3)
	fmt.Printf("%T", arr3)
}

func TestSlice(t *testing.T) {
	var slice = make([]int, 10)
	fmt.Printf("len: %d, cap: %d \n", len(slice), cap(slice))
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	var n = 3
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
		fmt.Println(dp[i][0])
	}
}
func TestSlice1(t *testing.T) {
	var sl1 []int
	sl1 = append(sl1, 1)
	fmt.Println(sl1[0])
}

func TestMap(t *testing.T) {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func TestMap1(t *testing.T) {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}

func modifyArray(arr *[3]int) {
	arr[0] = 100
}

type I interface {
	M()
}

type T struct {
}

func (T) M() {
}

func TestSwitch1(t1 *testing.T) {
	//var t T
	//var i I = t
	//switch i.(type) {
	//case T:
	//	println("it is type T")
	//case int:
	//	println("it is type int")
	//case string:
	//	println("it is type string")
	//}
}

func TestSwitch2(t *testing.T) {
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

type Book struct {
	Title   string         // 书名
	Pages   int            // 书的页数
	Indexes map[string]int // 书的索引
}

func (b Book) sayTitle() {
	fmt.Println(b.Title)
}

func TestEmpty(t *testing.T) {
	var book Book
	book.Title = "123"
	book.sayTitle()
	fmt.Println(book)
}
