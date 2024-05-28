package basic

import (
	"fmt"
	"testing"
)

// const 只能修饰基本变量，不能修饰结构体
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func TestConst(t *testing.T) {
	fmt.Println(KB)
}

func TestNumber(t *testing.T) {
	var number1 int32 = 30
	fmt.Printf("%T\n", number1)
	fmt.Printf("%T\n", int64(number1))
}

func TestString(t *testing.T) {
	traversalString()
	fmt.Println(modifyString("你好"))
}

func TestIfCase(t *testing.T) {
	var score = 100
	if score := 65; score >= 90 {
		fmt.Println("A", score)
	} else if score > 75 {
		fmt.Println("B", score)
	} else {
		fmt.Println("C", score)
	}
	// 注意变量覆盖
	fmt.Println(score)
}

func TestSwitch(t *testing.T) {
	switch s := "a"; s {
	case "a":
		fmt.Println("a")
		fallthrough
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

type NetError struct {
}

func (e NetError) Error() string {
	return "error"
}

func TestForceConvert(t *testing.T) {
	var bytes = []byte("123")
	fmt.Println(string(bytes))

}

func TestSum(t *testing.T) {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Printf("%T", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ {
		//byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, val := range s {
		//rune
		fmt.Printf("%v(%c) ", val, val)
	}
	fmt.Println()
}

func modifyString(str string) string {
	var strRune = []rune(str)
	strRune[0] = '刘'
	return string(strRune)
}
