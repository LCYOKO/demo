package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

func init() {
	fmt.Println("call init function()")
}

type person struct {
	name string
	age  int
}

//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
	age    int
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
	return r.width * r.height
}
func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (this *person) addAge(age int) {
	this.age += age
}

func newPerson() *person {
	return &person{
		name: "lisi",
		age:  12,
	}
}

func testStruct() {
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

func testToJson() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)
}

func testJsonToObj() {

}

func testError() {
	var myError error = nil
	fmt.Println(errors.New("123"))
	fmt.Println(myError)
}

func testNewPerson() {
	lisi := newPerson()
	fmt.Println(lisi)
	lisi.addAge(123)
	fmt.Println(lisi)
}

func testInterface() {
	var r geometry = &rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
