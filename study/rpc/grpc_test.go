package rpc

import (
	"context"
	"demo/pkg/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"testing"
)

type UserServer struct {
	pb.UnimplementedBookServiceServer
}

func (u *UserServer) GetBook(context context.Context, req *pb.BookRequest) (resp *pb.BookResponse, err error) {
	return &pb.BookResponse{Name: "golang"}, nil
}

func main() {

}

func TestGrpc(t *testing.T) {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 在gRPC服务端注册服务
	pb.RegisterBookServiceServer(s, &UserServer{})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
	var wg sync.WaitGroup
	wg.Wait()
}
