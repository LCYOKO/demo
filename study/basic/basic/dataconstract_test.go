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

	//	n2 := map[string]int{"foo": 1, "bar": 2}
	//	if maps.Equal(n, n2) {
	//		fmt.Println("n == n2")
	//	}
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
