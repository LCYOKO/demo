package grpc

//https://liwenzhou.com/posts/Go/gRPC/#c-0-7-7
import (
	"context"
	book "demo/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

type bookServer struct {
	book.UnimplementedBookServiceServer
}

func (b *bookServer) GetBooks1(context context.Context, req *book.BookRequest) (resp *book.BookResponse, err error) {
	md, ok := metadata.FromIncomingContext(context)
	if !ok {
		return nil, err
	}
	fmt.Printf("recive req:%T, md:%T", req, md)
	return generateResp(), nil
}
func (b *bookServer) GetBooks2(req *book.BookRequest, s book.BookService_GetBooks2Server) error {
	fmt.Printf("recive req:%T", req)
	s.Send(generateResp())
	time.Sleep(time.Second)
	s.Send(generateResp())
	time.Sleep(time.Second)
	s.Send(generateResp())
	return nil
}
func (b *bookServer) GetBooks3(s book.BookService_GetBooks3Server) error {
	resp := generateResp()
	for {
		// 接收客户端发来的流式数据
		res, err := s.Recv()
		if err == io.EOF {
			// 最终统一回复
			return s.SendAndClose(resp)
		}
		if err != nil {
			return err
		}
		fmt.Println(res)
	}
}
func (b *bookServer) GetBooks4(s book.BookService_GetBooks4Server) error {
	resp := generateResp()
	for {
		// 接收客户端发来的流式数据
		res, err := s.Recv()
		if err == io.EOF {
			// 最终统一回复
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(res)
		s.Send(resp)
	}
}

func generateResp() *book.BookResponse {
	books := []*book.Book{{
		Id:   1,
		Name: "golang",
	}}
	emptyMap := make(map[string]*book.Book, 0)
	return &book.BookResponse{Books: books, Status: book.Status_SUCCESS, BookMap: emptyMap}
}

func TestGrpc(t *testing.T) {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":6333")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	// 创建gRPC服务器
	server := grpc.NewServer()
	// 在gRPC服务端注册服务
	book.RegisterBookServiceServer(server, &bookServer{})
	var wg sync.WaitGroup
	wg.Add(1)
	// 启动服务
	go func() {
		err = server.Serve(lis)
		defer wg.Done()
		if err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
	}()
	wg.Wait()
}
