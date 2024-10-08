// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: book.proto

package book

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	GetBooks1(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	GetBooks2(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (BookService_GetBooks2Client, error)
	GetBooks3(ctx context.Context, opts ...grpc.CallOption) (BookService_GetBooks3Client, error)
	GetBooks4(ctx context.Context, opts ...grpc.CallOption) (BookService_GetBooks4Client, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBooks1(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/book.BookService/getBooks1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBooks2(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (BookService_GetBooks2Client, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[0], "/book.BookService/getBooks2", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceGetBooks2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_GetBooks2Client interface {
	Recv() (*BookResponse, error)
	grpc.ClientStream
}

type bookServiceGetBooks2Client struct {
	grpc.ClientStream
}

func (x *bookServiceGetBooks2Client) Recv() (*BookResponse, error) {
	m := new(BookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookServiceClient) GetBooks3(ctx context.Context, opts ...grpc.CallOption) (BookService_GetBooks3Client, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[1], "/book.BookService/getBooks3", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceGetBooks3Client{stream}
	return x, nil
}

type BookService_GetBooks3Client interface {
	Send(*BookRequest) error
	CloseAndRecv() (*BookResponse, error)
	grpc.ClientStream
}

type bookServiceGetBooks3Client struct {
	grpc.ClientStream
}

func (x *bookServiceGetBooks3Client) Send(m *BookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookServiceGetBooks3Client) CloseAndRecv() (*BookResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookServiceClient) GetBooks4(ctx context.Context, opts ...grpc.CallOption) (BookService_GetBooks4Client, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[2], "/book.BookService/getBooks4", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceGetBooks4Client{stream}
	return x, nil
}

type BookService_GetBooks4Client interface {
	Send(*BookRequest) error
	Recv() (*BookResponse, error)
	grpc.ClientStream
}

type bookServiceGetBooks4Client struct {
	grpc.ClientStream
}

func (x *bookServiceGetBooks4Client) Send(m *BookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookServiceGetBooks4Client) Recv() (*BookResponse, error) {
	m := new(BookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	GetBooks1(context.Context, *BookRequest) (*BookResponse, error)
	GetBooks2(*BookRequest, BookService_GetBooks2Server) error
	GetBooks3(BookService_GetBooks3Server) error
	GetBooks4(BookService_GetBooks4Server) error
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) GetBooks1(context.Context, *BookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks1 not implemented")
}
func (UnimplementedBookServiceServer) GetBooks2(*BookRequest, BookService_GetBooks2Server) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks2 not implemented")
}
func (UnimplementedBookServiceServer) GetBooks3(BookService_GetBooks3Server) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks3 not implemented")
}
func (UnimplementedBookServiceServer) GetBooks4(BookService_GetBooks4Server) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks4 not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_GetBooks1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBooks1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/getBooks1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBooks1(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBooks2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).GetBooks2(m, &bookServiceGetBooks2Server{stream})
}

type BookService_GetBooks2Server interface {
	Send(*BookResponse) error
	grpc.ServerStream
}

type bookServiceGetBooks2Server struct {
	grpc.ServerStream
}

func (x *bookServiceGetBooks2Server) Send(m *BookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BookService_GetBooks3_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookServiceServer).GetBooks3(&bookServiceGetBooks3Server{stream})
}

type BookService_GetBooks3Server interface {
	SendAndClose(*BookResponse) error
	Recv() (*BookRequest, error)
	grpc.ServerStream
}

type bookServiceGetBooks3Server struct {
	grpc.ServerStream
}

func (x *bookServiceGetBooks3Server) SendAndClose(m *BookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookServiceGetBooks3Server) Recv() (*BookRequest, error) {
	m := new(BookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BookService_GetBooks4_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookServiceServer).GetBooks4(&bookServiceGetBooks4Server{stream})
}

type BookService_GetBooks4Server interface {
	Send(*BookResponse) error
	Recv() (*BookRequest, error)
	grpc.ServerStream
}

type bookServiceGetBooks4Server struct {
	grpc.ServerStream
}

func (x *bookServiceGetBooks4Server) Send(m *BookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookServiceGetBooks4Server) Recv() (*BookRequest, error) {
	m := new(BookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getBooks1",
			Handler:    _BookService_GetBooks1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getBooks2",
			Handler:       _BookService_GetBooks2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getBooks3",
			Handler:       _BookService_GetBooks3_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getBooks4",
			Handler:       _BookService_GetBooks4_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "book.proto",
}
