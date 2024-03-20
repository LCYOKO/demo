package basic

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"testing"
)

func TestPrintf(t *testing.T) {
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	//%T	打印值的类型
	//%%	百分号
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	// b 二进制
	// c 字符
	// d 10进制整数
	// o 8进制
	// x 16进制
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
}

func TestSprint(t *testing.T) {
	s1 := fmt.Sprint("沙河小王子\n")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)
}

func TestErrorf(t *testing.T) {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误 \" %w \"", e)
	fmt.Println(w)
}

func TestScan(t *testing.T) {
	var (
		name    string
		age     int
		married bool
	)
	_, err := fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	if err != nil {
		fmt.Printf("Scan error %v", err)
		return
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func TestIoReader(t *testing.T) {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
