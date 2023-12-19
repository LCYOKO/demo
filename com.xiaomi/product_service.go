package com_xiaomi

import (
	"context"
	"fmt"
)

type BookService struct {
}

var MyProductService  = &BookService{}
func (this *BookService)GetBook(ctx context.Context, in *BookRequest) (*BookResponse, error){
	fmt.Println(in)
	return &BookResponse{Name:"123",Message:"success"}, nil
}