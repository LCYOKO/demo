package main

import (
	"context"
	com_xiaomi "demo/com.xiaomi"
	"demo/logger"
	"demo/redis"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)


type Config struct {
	Timeout string `json:"timeout"`
	Times int64
	Name string `json:"name"`
	Port int64 `json:"db.port"`
}

var config = new(Config)

func init()  {
 //viper.SetConfigFile("../config/config.yaml")
}

func main() {
 testGrpcClient()
	//testGrpcServer()
	//testViper()
	//Include(book.Route,user.Route)
	//engine := Init()
	//engine.Run(":9991")
	//fmt.Println(config)
}

func testGrpcServer()  {
	server := grpc.NewServer()
    com_xiaomi.RegisterBookServiceServer(server, com_xiaomi.MyProductService)
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	server.Serve(listen)
}

func testGrpcClient(){
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	prodClient := com_xiaomi.NewBookServiceClient(conn)
	prodRes, err := prodClient.GetBook(context.Background(),
		&com_xiaomi.BookRequest{Id: 12,Name:"123"})
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes)

}

func testViper()  {
	viper.SetConfigFile("./config/config.yaml")
	err:=viper.ReadInConfig()
	if err!=nil {
		panic(fmt.Errorf("Fatal read config failed: %s \n", err))
	}
	if err:=viper.Unmarshal(config);err!=nil {
		panic(fmt.Errorf("Fatal read config failed: %s \n", err))
	}
	fmt.Println(viper.Get("name"))
	fmt.Println(config)
	fmt.Println(viper.Get("db.port"))
}

func testLogger(){
logger.InitLogger(logger.LogConfig{
	Level:      "info",
	Filename:   "./demo.log",
	MaxSize:    10000,
	MaxAge:     3600,
	MaxBackups: 0,
})
}

func testZookeeper(){

}

func testKafka()(){

}

func testRabbitMq(){

}

func testRedis(){
	error := redis.Init()
	if error!=nil {
		fmt.Println(error)
		os.Exit(0)
	}
	testKey := "testStr"
	redis.RedisClient.Set(testKey,"123",-1)
	stringCmd := redis.RedisClient.Get(testKey)
	fmt.Println(stringCmd.Result())
}

