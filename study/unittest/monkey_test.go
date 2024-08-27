package unittest

import (
	"bou.ke/monkey"
	"demo/study/data"
	"fmt"
	"github.com/magiconair/properties/assert"
	"os"
	"strings"
	"testing"
)

func MyFunc(uid int64) string {
	u, err := data.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return u.Name
}

// go test -run=TestMyFunc -v -gcflags=-l 注意需要避免内联优化
func TestMyFunc(t *testing.T) {
	// 对GetInfoByUID 进行打桩
	monkey.Patch(data.GetInfoByUID, func(int64) (*data.UserInfo, error) {
		return &data.UserInfo{Name: "liwenzhou"}, nil
	})

	ret := MyFunc(123)
	assert.Equal(t, ret, "liwenzhou")
}

func Test2(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?")
}
