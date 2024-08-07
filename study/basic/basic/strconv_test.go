package basic

import (
	"fmt"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	str := "123"
	parseInt, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("str:%s, parseInt:%d, type:%T \n", str, parseInt, parseInt)
}

func Test2(t *testing.T) {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println(s1, s2, s3, s4)
}
