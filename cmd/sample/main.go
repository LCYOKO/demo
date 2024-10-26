package main

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
	if err != nil {
		log.Fatal("store factor init failed", err)
		return
	}
	srv := server.NewServer(":8080", s)
	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("web server start failed", err)
		return
	}
	log.Println("web server start success")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err = <-errChan:
		log.Fatal("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("bookstore program exit failed", err)
		return
	}
	log.Println("bookstore program exit success")
}
