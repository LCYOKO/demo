package _struct

import (
	"fmt"
	"testing"
)

type parent struct {
	Name string
}

func (p parent) name() string {
	return p.Name
}

func (p parent) say() {
	fmt.Printf("hello %s", p.name())
}

type son struct {
	parent
	Name string
}

func (s son) name() string {
	return s.name()
}

//func (s son) say()  {
//	fmt.Printf("hello %s", s.name())
//}

func TestRewrite(t *testing.T) {
	//p := &parent{
	//	Name: "parent",
	//}

	s := &son{
		Name: "son",
	}
	s.say()
}
