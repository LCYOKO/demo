package filters

import (
	"demo/study/practice"
	"fmt"
)

func init() {
	practice.RegisterFilter("my-custom", myFilterBuilder)
}

func myFilterBuilder(next practice.Filter) practice.Filter {
	return func(c *practice.Context) {
		fmt.Println("假装这是我自定义的 filter")
		next(c)
	}
}
