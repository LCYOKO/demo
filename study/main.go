package main

import (
	"fmt"
	"maps"
)

func main() {
	//基础数据类型
	//testMap
	//testSum()
	//testRune()
	//testRune()

	//基本操作fmt
	//testSFmt()
	//testEFmt()

	//文件操作
	test
	//结构体相关
	//testNewPerson()
	//testStruct()
	//testInterface()
	//testToJson()
	//testReflectType()

	//并发
	//testGoro()
	//testChannel1()
	//testChannel2()
	//testChannel3()
	//testChannel4()
	//testOnce()
}

func testRune() {
	cnName := []rune("你好")
	fmt.Println(len(cnName))
	for _, word := range cnName {
		fmt.Printf("%c\n", word)
	}
}

func testMap() {
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

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func testSum() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
