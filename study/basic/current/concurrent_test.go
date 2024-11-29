package current

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func TestChannel1(t *testing.T) {
	ch := make(chan int)
	fmt.Println(ch)
	go func() {
		ch <- 12
	}()
	fmt.Println("recvie", <-ch)
}

func TestChannel2(t *testing.T) {
	ch := make(chan int)
	<-ch
	fmt.Println("123")
	//对一个关闭的通道再发送值就会导致 panic。
	//对一个关闭的通道进行接收会一直获取值直到通道为空。
	//对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	//关闭一个已经关闭的通道会导致 panic。
	close(ch)
}

func TestChannel3(t *testing.T) {
	ch := make(chan int)
	go func(ch chan int) {
		for val := range ch {
			fmt.Println(val)
		}
		fmt.Println("break")
	}(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func TestChannel4(t *testing.T) {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func TestChannelCancel(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			select {
			case val := <-ch:
				fmt.Println("val", val)
				break
			}
		}(i)
	}
	time.Sleep(time.Second)
	ch <- 10
	close(ch)
	wg.Wait()
}

// 很经典的死循环
func TestChannelError1(t *testing.T) {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				//无限读0值
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task)
			}
			//FIXME
			wg.Done()
		}()
	}
	wg.Wait()
}

// 由于 select 命中了超时逻辑，导致通道没有消费者（无接收操作），而其定义的通道为无缓冲通道
// 因此 goroutine 中的ch <- "job result"操作会一直阻塞，最终导致 goroutine 泄露。
func TestChannelError2(t *testing.T) {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}

var (
	x       int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)
}

func TestMutex(t *testing.T) {
	// 使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 1000) // x:10 cost:1.466500951s
	// 使用读写互斥锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms
}

func TestOnce(t *testing.T) {
	num := 10
	group := sync.WaitGroup{}
	group.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer group.Done()
			getInstance := GetInstance()
			fmt.Printf("%p\n", getInstance)
		}()
	}
	group.Wait()
}

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func TestCurrent1(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 1; i < 10; i++ {
		go func() {
			//FIXME
			fmt.Println("A", i)
		}()
	}
	time.Sleep(time.Hour)
}

func TestErrorGroup(t *testing.T) {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	if err := g.Wait(); err != nil {
		// 处理可能出现的错误
		fmt.Println(err)
		return
	}
	fmt.Println("所有goroutine均成功")
}

//func GetFriends(ctx context.Context, user int64) (map[string]*User, error) {
//	g, ctx := errgroup.WithContext(ctx)
//	friendIds := make(chan int64)
//
//	// Produce
//	g.Go(func() error {
//		defer close(friendIds)
//		for it := GetFriendIds(user); ; {
//			if id, err := it.Next(ctx); err != nil {
//				if err == io.EOF {
//					return nil
//				}
//				return fmt.Errorf("GetFriendIds %d: %s", user, err)
//			} else {
//				select {
//				case <-ctx.Done():
//					return ctx.Err()
//				case friendIds <- id:
//				}
//			}
//		}
//	})
//
//	friends := make(chan *User)
//
//	// Map
//	workers := int32(nWorkers)
//	for i := 0; i < nWorkers; i++ {
//		g.Go(func() error {
//			defer func() {
//				// Last one out closes shop
//				if atomic.AddInt32(&workers, -1) == 0 {
//					close(friends)
//				}
//			}()
//
//			for id := range friendIds {
//				if friend, err := GetUserProfile(ctx, id); err != nil {
//					return fmt.Errorf("GetUserProfile %d: %s", user, err)
//				} else {
//					select {
//					case <-ctx.Done():
//						return ctx.Err()
//					case friends <- friend:
//					}
//				}
//			}
//			return nil
//		})
//	}
//	// Reduce
//	ret := map[string]*User{}
//	g.Go(func() error {
//		for friend := range friends {
//			ret[friend.Name] = friend
//		}
//		return nil
//	})
//	return ret, g.Wait()
//}

func TestSync(t *testing.T) {
	//同步类型的值不能被拷贝,以下类型均不能被拷贝
	//sync.Cond
	//sync.Map
	//sync.Mutex
	//sync.RWMutex
	//sync.Once
	//sync.Pool
	//sync.WaitGroup
}
