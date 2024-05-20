package main

import (
	"demo/internal/web"
	"os"
)

func main() {
	command := web.NewCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

//func InitGin() {
//	routers2.Include(book2.Routers, user2.Routers)
//	engine := routers2.Init()
//	err := engine.Run("localhost:8081")
//	if err != nil {
//		return
//	}
//}
