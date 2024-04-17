package main

import (
	"demo/internal"
	book2 "demo/internal/controller/book"
	user2 "demo/internal/controller/user"
	routers2 "demo/internal/routers"
	"os"
)

type Config struct {
	Timeout string `json:"timeout"`
	Times   int64
	Name    string `json:"name"`
	Port    int64  `json:"db.port"`
}

var config = new(Config)

func main() {
	command := internal.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func InitGin() {
	routers2.Include(book2.Routers, user2.Routers)
	engine := routers2.Init()
	err := engine.Run("localhost:8081")
	if err != nil {
		return
	}
}
