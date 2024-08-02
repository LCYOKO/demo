package main

// package
//  hello
import (
	"context"
	"demo/internal/sample/server"
	"demo/internal/sample/store/factor"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factor.New("memo")
	log.Println("s", s)
	srv := server.NewServer(":8080", s)

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("web server start failed:", err)
		return
	}
	log.Println("web server start ok")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select { // 监视来自errChan以及c的事件
	case err = <-errChan:
		log.Fatal("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx) // 优雅关闭http服务实例
	}

	if err != nil {
		log.Println("bookstore program exit error:", err)
		return
	}
	log.Println("bookstore program exit ok")
}
