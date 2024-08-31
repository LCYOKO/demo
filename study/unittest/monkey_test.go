package unittest

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/magiconair/properties/assert"
	"os"
	"strings"
	"testing"
)

type UserInfo struct {
	Name string
}

func GetInfoByUID(uid int64) (*UserInfo, error) {
	return &UserInfo{
		Name: "lisi",
	}, nil
}
func MyFunc(uid int64) string {
	u, err := GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return u.Name
}

// go test -run=TestMyFunc -v -gcflags=-l 注意需要避免内联优化
func TestMyFunc(t *testing.T) {
	// 对GetInfoByUID 进行打桩
	monkey.Patch(GetInfoByUID, func(int64) (*UserInfo, error) {
		return &UserInfo{Name: "liwenzhou"}, nil
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
