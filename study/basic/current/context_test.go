package current

import (
	"context"
	"fmt"
	"testing"
	"time"
)

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
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancelFunc()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestDeadLine(t *testing.T) {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*5))
	defer cancelFunc()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestWithValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "traceId", "v1")
	test1(ctx)
}

func test1(c context.Context) {
	fmt.Println(c.Value("traceId"))
}
