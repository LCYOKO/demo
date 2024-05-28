package _struct

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerType(t *testing.T) {
	type MyInt int
	var numI = 123
	var num MyInt = MyInt(numI)
	// 自定义类型是新类型
	fmt.Printf("type: %T", num)
}

func TestTypeAlias(t *testing.T) {
	type MyInt = int
	var numI = 123
	var num MyInt = numI
	//别名和新类型一样，只是名称不同
	fmt.Printf("%T", num)
}

type People struct {
	name string //私有不能被json包访问
	age  int
}

//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
	age    int
}

func TestStruct1(t *testing.T) {
	var stuTest Student
	fmt.Println("空值", stuTest)
	var stu = &Student{
		ID:     1,
		Gender: "男",
	}
	fmt.Println(stu)

	m := make(map[string]*Student)
	stus := []Student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func TestJson(t *testing.T) {
	var stu = &Student{
		ID:     1,
		Gender: "男",
	}
	jsonStr, err := json.Marshal(stu)
	if err != nil {
		return
	}
	fmt.Println(string(jsonStr))
}
