package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestError(t *testing.T) {
  fmt.Println(throw())
}

func throw() error {
	return errors.New("error")
}
