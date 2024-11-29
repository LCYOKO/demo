package context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			fmt.Println("worker over")
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			fmt.Println("worker2 over")
			break LOOP
		default:
		}
	}
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

func TestTimeout(t *testing.T) {
	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*5)
	//defer cancelFunc()
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err())
	//}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()
	// 输出2，说明在 ctx.Done()这里阻塞了两秒
	fmt.Println(end - start)
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		fmt.Println("context was canceled")
	}()
	// 确保我们的 goroutine进去执行了
	time.Sleep(time.Second)
	cancel()
	// 确保后面那句打印出来了
	time.Sleep(time.Second)
}

func TestDeadLine(t *testing.T) {
	//ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*5))
	//defer cancelFunc()
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err())
	//}
	// 设置两秒后超时
	ctx, cancel := context.WithDeadline(context.Background(),
		time.Now().Add(2*time.Second))
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()
	// 输出2，说明在 ctx.Done()这里阻塞了两秒
	fmt.Println(end - start)

}

func TestWithValue(t *testing.T) {
	parentKey := "parent"
	parent := context.WithValue(context.Background(), parentKey, "this is parent")

	sonKey := "son"
	son := context.WithValue(parent, sonKey, "this is son")

	// 尝试从 parent 里面拿出来 key = son的，会拿不到
	if parent.Value(parentKey) == nil {
		fmt.Printf("parent can not get son's key-value pair")
	}

	if val := son.Value(parentKey); val != nil {
		fmt.Printf("parent can not get son's key-value pair")
	}
}
