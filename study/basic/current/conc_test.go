package current
//https://www.liwenzhou.com/posts/Go/conc/#c-0-4-2
//https://github.com/sourcegraph/conc
import (
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/stream"
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup1(t *testing.T) {
	var count atomic.Int64
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			count.Add(1)
		})
	}
	// 等待10个goroutine都执行完
	wg.Wait()
	fmt.Println(count.Load())

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			count.Add(1)
		})
	}
	// 等待10个goroutine都执行完
	wg.Wait()
	fmt.Println(count.Load())
}

func TestWaitGroup2(t *testing.T) {
	var count atomic.Int64
	var wg conc.WaitGroup
	// 开启10个goroutine并发执行 count.Add(1)
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			if i == 7 {
				panic("bad thing")
			}
			count.Add(1)
		})
	}
	// 等待10个goroutine都执行完
	wg.WaitAndRecover()
	fmt.Println(count.Load())
}

func TestStream(t *testing.T) {
	//使用Stream时提交的每个任务都返回一个回调函数。每个任务都将在任务池中同时执行，但是回调函数将按照任务提交的顺序依次执行。
	//等到所有任务都提交后，必须调用 Wait()来等待正在运行的 goroutine 运行完。当任务执行过程中或回调函数执行期间出现 panic 时
	//，所有其他任务和回调仍将执行。当调用 Wait()时，panic将传给调用方。
	//同Pool一样，Stream也不适用于非常短的任务。启动和拆卸会增加几微秒的开销，每个任务的开销大约是500ns。对于任何需要网络通话的任务来说，这性能都足够好了
	times := []int{20, 52, 16, 45, 4, 80}
	s := stream.New()
	for _, millis := range times {
		dur := time.Duration(millis) * time.Millisecond
		// 提交任务
		s.Go(func() stream.Callback {
			time.Sleep(dur)
			// 虽然上一行通过sleep增加了时间
			// 但最终结果仍按任务提交（s.Go）的顺序打印
			return func() { fmt.Println(dur) }
		})
	}
	s.Wait()
}

func TestIterator(t *testing.T) {
	//input := []int{1, 2, 3, 4}
	//// 创建一个最大goroutine个数为输入元素一半的迭代器
	//iterator := iter.Iterator[int]{
	//	MaxGoroutines: len(input) / 2,
	//}
	//
	//iterator.ForEach(input, func(v *int) {
	//	if *v%2 != 0 {
	//		*v = -1
	//	}
	//})
	//fmt.Println(input)
}

func TestMapper(t *testing.T) {
	//input := []int{1, 2, 3, 4}
	//// 创建一个最大goroutine个数为输入元素一半的映射器
	//mapper := iter.Mapper[int, bool]{
	//	MaxGoroutines: len(input) / 2,
	//}
	//
	//results := mapper.Map(input, func(v *int) bool { return *v%2 == 0 })
	//fmt.Println(results)
}
