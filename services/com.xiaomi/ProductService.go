package com_xiaomi

import (
	"context"
	grpc "google.golang.org/grpc"
)

type server struct {
}

func (this *server) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error){
	 	return &ProdResponse{ProdStock:20},nil
}