package main

import (
	"demo/internal/miniblog"
	"os"
)

func main() {
	err := miniblog.NewMiniBlogCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
