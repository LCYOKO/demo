package goin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

/**
https://gin-gonic.com/zh-cn/docs/introduction/

gin.Default()默认使用了Logger和Recovery中间件，其中：
Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

gin中间件中使用goroutine
当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
*/
var engine *gin.Engine

type Login struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required" message:"password should not empty" `
}

func setup() {
	engine = gin.Default()
}

func TestMain(m *testing.M) {
	setup()
	run := m.Run()
	os.Exit(run)
}

func TestQuery(t *testing.T) {
	engine.GET("/get", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	engine.POST("/post", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	engine.POST("/json", func(c *gin.Context) {
		data, _ := c.GetRawData()
		fmt.Println(string(data))
		var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(data, &m)

		c.JSON(http.StatusOK, m)
	})

	engine.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	engine.Run(":8080")
}

func TestBind(t *testing.T) {
	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	engine.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	engine.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	engine.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	engine.Run(":8080")
}

func TestSingleUpload(t *testing.T) {
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	engine.POST("/uploadSingle", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		fmt.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	engine.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			fmt.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	engine.Run(":8080")
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
		log.Fatal("server shutdown: .", err)
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
