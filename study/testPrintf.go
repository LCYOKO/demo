package main

import (
	"errors"
	"fmt"
)

func testFmt() {

	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	//%T	打印值的类型
	//%%	百分号
}

func testSFmt() {
	s1 := fmt.Sprint("沙河小王子\n")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)
}

func testEFmt() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误 \" %w \"", e)
	fmt.Println(w)
}
