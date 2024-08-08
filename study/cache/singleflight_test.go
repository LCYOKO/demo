package cache

//https://www.liwenzhou.com/posts/Go/singleflight/
import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"testing"
	"time"
)

func getData(val int64) int64 {
	fmt.Println("start get data")
	time.Sleep(5 * time.Second)
	return val + 1
}

func Test1(t *testing.T) {
	g := new(singleflight.Group)

	// 第1次调用
	go func() {
		v1, _, shared := g.Do("getData", func() (interface{}, error) {
			ret := getData(1)
			return ret, nil
		})
		fmt.Printf("1st call: v1:%v, shared:%v\n", v1, shared)
	}()

	time.Sleep(2 * time.Second)

	// 第2次调用（第1次调用已开始但未结束）
	v2, _, shared := g.Do("getData", func() (interface{}, error) {
		ret := getData(1)
		return ret, nil
	})
	fmt.Printf("2nd call: v2:%v, shared:%v\n", v2, shared)
}

func doChanGetData(ctx context.Context, g *singleflight.Group, id int64) (int64, error) {
	ch := g.DoChan("getData", func() (interface{}, error) {
		ret := getData(id)
		return ret, nil
	})
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case ret := <-ch:
		return ret.Val.(int64), ret.Err
	}
}

func Test2(t *testing.T) {
	g := &singleflight.Group{}

	// 第1次调用
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		v1, err := doChanGetData(ctx, g, 1)
		fmt.Printf("v1:%v err:%v\n", v1, err)
	}()

	time.Sleep(2 * time.Second)

	// 第2次调用（第1次调用已开始但未结束）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	v2, err := doChanGetData(ctx, g, 1)
	fmt.Printf("v2:%v err:%v\n", v2, err)
}

func doGetData(g *singleflight.Group, id int64) (string, error) {
	v, err, _ := g.Do("getData", func() (interface{}, error) {
		go func() {
			time.Sleep(100 * time.Millisecond) // 100ms后忘记key
			g.Forget("getData")
		}()

		ret := getData(id)
		return ret, nil
	})
	return v.(string), err
}
