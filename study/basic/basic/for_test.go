package basic

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	arr := []int{9, 8, 7, 6}

	for index, value := range arr {
		fmt.Printf("%d => %d\n", index, value)
	}

	// 如果只是需要 value, 可以用 _ 代替 index
	for _, value := range arr {
		fmt.Printf("only value: %d \n", value)
	}

	// 如果只需要 index 也可以去掉 写成 for index := range arr
	for index := range arr {
		fmt.Printf("only index: %d \n", index)
	}

	fmt.Println("for r loop end \n ")
}
