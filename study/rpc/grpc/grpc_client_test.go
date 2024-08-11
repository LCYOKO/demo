package grpc

import (
	"context"
	book "demo/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

var c book.BookServiceClient

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:6333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = book.NewBookServiceClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBooks1(ctx, &book.BookRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetBooks())
}

func TestClient2(t *testing.T) {
	conn, err := grpc.Dial("localhost:6333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = book.NewBookServiceClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := c.GetBooks2(ctx, &book.BookRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		recv, err := r.Recv()
		if err != nil {
			fmt.Printf("error %+v", err)
			return
		}
		fmt.Println(recv.Books)
	}
}

func TestClient3(t *testing.T) {
	conn, err := grpc.Dial("localhost:6333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = book.NewBookServiceClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := c.GetBooks3(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	r.Send(&book.BookRequest{Id: 1})
	r.Send(&book.BookRequest{Id: 1})
	r.Send(&book.BookRequest{Id: 1})
	recv, err := r.CloseAndRecv()
	if err != nil {
		return
	}
	fmt.Println(recv)
}

func TestClient4(t *testing.T) {
	conn, err := grpc.Dial("localhost:6333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = book.NewBookServiceClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := c.GetBooks4(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	r.Send(&book.BookRequest{Id: 1})
	r.Send(&book.BookRequest{Id: 1})
	r.Send(&book.BookRequest{Id: 1})
	r.CloseSend()
	for {
		recv, err := r.Recv()
		if err != nil {
			fmt.Printf("error %v\n", err)
			return
		}
		fmt.Printf("recv %v \n", recv)
	}
}
