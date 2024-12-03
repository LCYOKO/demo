package concurrent

import (
	"context"
	"fmt"
	"go.uber.org/atomic"
	"testing"
	"time"
)

const _CLOSE int32 = 0
const _STARTED int32 = 1

type CacheBlockTaskPool struct {
	concurrent chan struct{}
	queue      chan func()
	cnt        atomic.Int32
	ctx        context.Context
	cancel     context.CancelFunc
	// 0关闭  1启动
	state atomic.Int32
}

func NewCacheBlockTaskPool(maxCon int, queueSize int) *CacheBlockTaskPool {
	ctx, cancelFunc := context.WithCancel(context.Background())
	tp := &CacheBlockTaskPool{
		concurrent: make(chan struct{}, maxCon),
		queue:      make(chan func(), queueSize),
		ctx:        ctx,
		cancel:     cancelFunc,
	}
	return tp
}

func (tp *CacheBlockTaskPool) Do(action func()) {
	if tp.IsStarted() {
		tp.queue <- action
	}
}

func (tp *CacheBlockTaskPool) Start() {
	if ok := tp.state.CAS(_CLOSE, _STARTED); ok {
		go func() {
			for {
				select {
				case tp.concurrent <- struct{}{}:
					tp.run()
				case <-tp.ctx.Done():
					fmt.Println("routine pool down ")
					close(tp.concurrent)
					close(tp.queue)
					return
				}
			}
		}()
	}
}

func (tp *CacheBlockTaskPool) run() {
	go func(ctx context.Context) {
		tp.cnt.Add(1)
		for {
			isDown := false
			//fmt.Printf("task goroutine num %d , queue len: %d \n", tp.cnt.Load(), len(tp.queue))
			select {
			case <-ctx.Done():
				isDown = true
				break
			case task := <-tp.queue:
				if task == nil {
					isDown = true
					break
				}
				task()
			}
			if isDown {
				break
			}
		}
		<-tp.concurrent
		tp.cnt.Sub(1)
	}(context.WithoutCancel(tp.ctx))
}

func (tp *CacheBlockTaskPool) IsStarted() bool {
	return tp.state.Load() == _STARTED
}
func (tp *CacheBlockTaskPool) Close() {
	if tp.IsStarted() {
		if ok := tp.state.CAS(_STARTED, _CLOSE); ok {
			tp.cancel()
		}
	}
}

func (tp *CacheBlockTaskPool) printStat() {
	fmt.Printf("routineCnt:%v, queuenSize:%d, status:%v \n", tp.cnt, len(tp.queue), tp.state.Load())
}

func TestRoutinePool(t *testing.T) {
	pool := NewCacheBlockTaskPool(10, 50)
	go func() {
		for {
			if pool.IsStarted() && len(pool.queue) == 0 {
				pool.Close()
			}
			pool.printStat()
			time.Sleep(time.Millisecond * 100)
		}
	}()
	pool.Start()

	for i := 0; i < 100; i++ {
		taskId := i
		pool.Do(func() {
			fmt.Printf("call task:%d \n", taskId)
			time.Sleep(time.Second)
		})
	}
	time.Sleep(time.Hour * 1)
}
