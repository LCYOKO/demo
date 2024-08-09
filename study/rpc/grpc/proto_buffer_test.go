package grpc

import (
	book "demo/pb"
	"fmt"
	"testing"
)

func TestProto(t *testing.T) {
	book := book.Book{Id: 1, Name: "Java实战"}
	fmt.Println(book.Name)
}
