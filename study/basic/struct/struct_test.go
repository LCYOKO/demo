package _struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
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

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

func TestMethod(t1 *testing.T) {
	var t T
	f1 := (*T).Set                           // f1的类型，也是*T类型Set方法的类型：func (t *T, int)int
	f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T,
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T)
	f1(&t, 3)
	fmt.Println(f2(t)) // 3
}

type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (*T1) PT1M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (*T2) PT2M2() { println("PT2's M2") }

type T3 struct {
	T1
	*T2
}

func TestInherit(t1 *testing.T) {
	//类型 T 的方法集合 = T1 的方法集合 + *T2 的方法集合
	//类型 *T 的方法集合 = *T1 的方法集合 + *T2 的方法集合
	t := T3{
		T1: T1{},
		T2: &T2{},
	}
	dumpMethodSet(t)
	dumpMethodSet(&t)
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func TestInvalidUseCase(t *testing.T) {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}
	//FIXME 看下如何把结果改对
	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}

// Student 学生
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

type UserSvc interface {
	sayHello()
	sayName()
}

type AbsSvc struct {
	Name string
}

func (a *AbsSvc) sayHello() {
	fmt.Println("sayHello")
}
func (a *AbsSvc) sayName() {
	fmt.Println(a.Name)
}

type UserSvcImpl struct {
	*AbsSvc
}

func (u *UserSvcImpl) sayHello() {
	fmt.Println("override sayHello")
}

func TestAbs(t *testing.T) {
	var u UserSvc = &UserSvcImpl{&AbsSvc{Name: "abs"}}
	u.sayName()
	u.sayHello()
}
