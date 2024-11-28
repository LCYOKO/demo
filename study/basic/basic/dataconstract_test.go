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

	//修改数组
	modifyArray(&arr3)
	fmt.Println(arr3)
	fmt.Printf("%T", arr3)
}

func TestSlice(t *testing.T) {
	var slice = make([]int, 10)
	fmt.Printf("len: %d, cap: %d \n", len(slice), cap(slice))
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		// 注意append操作不是线程安全的
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

	sl1 = make([]int, 10)
	sl2 := sl1[1:3]
	sl1[3] = 4
	fmt.Println(sl1)
	_ = append(sl2, 1)
	fmt.Println(len(sl2), cap(sl2), sl1)
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

func TestCompare1(t *testing.T) {
	//分类	说明	                         是否能比较	                     说明
	//基本类型	整型（ int/uint/int8/uint8/int16/uint16/int32/uint32/int64/uint64/byte/rune等）浮点数（ float32/float64）复数类型（ complex64/complex128）字符串（ string）	是
	//引用类型	切片（slice）、map	      否
	//channel、指针	                      是
	//聚合类型（复合类型）	数组	              是	              相同长度的数组可以比较，不同长度的数组不能进行比较
	//结构体	                              是	               只包含可比较的类型情况下可比较
	//接口类型	如error	                  是

	//引用类型
	//slice、map
	//切片之间不允许比较。切片只能与nil值比较
	//map之间不允许比较。map只能与nil值比较
	//两个nil也不能比较，会panic
	//slice、map比较

	//使用reflect.DeepEqual()
	//对比规则
	//相同类型的值是深度相等的，不同类型的值永远不会深度相等。
	//当数组值（array）的对应元素深度相等时，数组值是深度相等的。
	//当结构体（struct）值如果其对应的字段（包括导出和未导出的字段）都是深度相等的，则该值是深度相等的。
	//当函数（func）值如果都是零，则是深度相等；否则就不是深度相等。
	//当接口（interface）值如果持有深度相等的具体值，则深度相等。
	//当切片（slice）序号相同，如果值,指针都相等，那么就是深度相等的
	//当哈希表（map）相同的key，如果值，指针都相等，那么就是深度相等的。

	//channel、指针
	//指针可比较，只要指针指向的地址一样，则相等
	//由于通过make创建channel后，返回的是一个指针，所以可以比较
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
