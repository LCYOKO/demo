package main

import (
	"context"
	"demo/common"
	"demo/controller/book"
	"demo/controller/user"
	testgorm "demo/gorm"
	"demo/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Timeout string `json:"timeout"`
	Times   int64
	Name    string `json:"name"`
	Port    int64  `json:"db.port"`
}

var config = new(Config)

/**
https://gin-gonic.com/zh-cn/docs/introduction/

gin.Default()默认使用了Logger和Recovery中间件，其中：
Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

gin中间件中使用goroutine
当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
*/
func main() {
	//testDefaultGin()
	//testGinGroup()
	//testMiddleWare()
	testgorm.TestCreate()
	//InitGin()
}

func InitGin() {
	routers.Include(book.Routers, user.Route)
	engine := routers.Init()
	engine.Run("localhost:8081")
}

func testDefaultGin() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, common.Ok(gin.H{"message": "Hello world!"}))
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, common.Ok(msg))
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}

func testGinGroup() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.Ok("/user/index"))
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.Ok("/user/index"))
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.Ok("/user/index"))
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.Ok("/shop/index"))
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.Ok("/shop/cart"))
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.Ok("/shop/checkout"))

		})
	}
	r.Run()
}

func testMiddleWare() {
	StatCost := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			start := time.Now()
			// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
			c.Set("name", "小王子")
			// 调用该请求的剩余处理程序
			c.Next()
			// 不调用该请求的剩余处理程序
			// c.Abort()
			// 计算耗时
			cost := time.Since(start)
			log.Println(cost)
		}
	}
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(StatCost())
	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}

func testGracefulShutdown() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

func testGracefulRestart() {
	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "hello gin!")
	//})
	//// 默认endless服务器会监听下列信号：
	//// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	//// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	//// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	//// 接收到 SIGUSR2 信号将触发HammerTime
	//// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	//if err := endless.ListenAndServe(":8080", router); err != nil {
	//	log.Fatalf("listen: %s\n", err)
	//}
	//log.Println("Server exiting")
}
