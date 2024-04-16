package unittest

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	// 程序输出的结果
	got := Split("a:b:c", ":")
	// 期望的结果
	want := []string{"a", "b", "c"}
	// 因为slice不能比较直接，借助反射包中的方法比较
	assert.Equal(t, got, want)
}

func TestSplitWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	assert.Equal(t, got, want)
}

func TestSplitAll(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tt.input, tt.sep)
			assert.Equal(t, got, tt.want)
		})
	}
}

//
//func testGocovery() {
//	// 只需要在顶层的Convey调用时传入t
//	c.Convey("分隔符在开头或结尾用例", t, func() {
//		tt := []struct {
//			name   string
//			s      string
//			sep    string
//			expect []string
//		}{
//			{"分隔符在开头", "*1*2*3", "*", []string{"", "1", "2", "3"}},
//			{"分隔符在结尾", "1+2+3+", "+", []string{"1", "2", "3", ""}},
//		}
//		for _, tc := range tt {
//			c.Convey(tc.name, func() { // 嵌套调用Convey
//				got := Split(tc.s, tc.sep)
//				c.So(got, c.ShouldResemble, tc.expect)
//			})
//		}
//	})
//}

func testMonkey() {
	//list.New().
	//monkey 好用的打桩工具
}
