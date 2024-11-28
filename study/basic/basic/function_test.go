package basic

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func add(x, y int) int {
	return x + y
}

func cacl(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, cacl(1, 2, add))
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func TestFunc1(t *testing.T) {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

func TestDefer1(t *testing.T) {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func TestDefer2(t *testing.T) {
	// 5
	fmt.Println(f1())
	// 6
	fmt.Println(f2())
	// 5
	fmt.Println(f3())
	// 5
	fmt.Println(f4())
}

func calc1(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func TestDefer3(t *testing.T) {
	x := 1
	y := 2
	defer calc1("AA", x, calc1("A", x, y))
	x = 10
	defer calc1("BB", x, calc1("B", x, y))
	y = 20
	// A, 1, 2, 3
	// B, 10, 2, 12
	// BB, 10, 12, 22
	// AA, 1, 3, 4
}

func TestPanic(t *testing.T) {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B", err)
		}
	}()
	panic("panic in B")
}

func TestFun1(t *testing.T) {
	sl := make([]int, 0, 10)
	fmt.Println(sl)
	sl = myAppend(sl, 1, 2, 3, 4)
	fmt.Println(sl)
}

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems)
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}
	sl = append(sl, elems...)
	return sl
}
