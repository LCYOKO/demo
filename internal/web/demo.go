package web

import (
	"context"
	conf2 "demo/internal/web/conf"
	controller2 "demo/internal/web/controller"
	logger2 "demo/internal/web/logger"
	store2 "demo/internal/web/store"
	middleware "demo/pkg/middware"
	"demo/pkg/version/verflag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfgFile string

// NewCommand 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "web",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.
Find more miniblog information at:
	https://github.com/marmotedu/miniblog#readme`,
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	//cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。
	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	if err := conf2.Init(); err != nil {
		return err
	}
	if err := logger2.Init(&conf2.Conf.Log); err != nil {
		return err
	}
	//	// 设置 token 包的签发密钥，用于 token 包 token 的签发和解析
	//	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)
	var g *gin.Engine
	var err error
	if g, err = initGin(conf2.Conf); err != nil {
		return err
	}
	// 初始化 store 层
	if err := store2.Init(conf2.Conf); err != nil {
		return err
	}
	// 初始化service

	// 初始化路由
	if err := controller2.Init(g); err != nil {
		return err
	}
	// 创建并运行 HTTP 服务器
	httpSrv := startInsecureServer(g, &conf2.Conf.Http)
	//
	//	// 创建并运行 HTTPS 服务器
	//	httpssrv := startSecureServer(g)
	//
	//	// 创建并运行 GRPC 服务器
	grpcSrv := startGRPCServer(conf2.Conf)
	return graceFullShutDown(httpSrv, grpcSrv)
}

func initGin(conf *conf2.Config) (g *gin.Engine, err error) {
	//设置 Gin 模式
	logger2.Logger.Info("config", zap.String("model", conf.Model))
	gin.SetMode(conf.Model)
	// 创建 Gin 引擎
	g = gin.New()
	g.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	return g, err
}

func graceFullShutDown(httpSrv *http.Server, grpcSrv *grpc.Server) error {
	var log = logger2.SugaredLogger
	// 等待中断信号优雅地关闭服务器（10 秒超时)。
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Info("Shutting down server ...")

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Error("Insecure Server forced to shutdown", zap.Error(err))
		return err
	}
	//	if err := httpssrv.Shutdown(ctx); err != nil {
	//		log.Errorw("Secure Server forced to shutdown", "err", err)
	//		return err
	//	}
	//
	if grpcSrv != nil {
		grpcSrv.GracefulStop()
	}
	log.Info("Server exiting")
	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine, conf *conf2.Http) *http.Server {
	// 创建 HTTP Server 实例
	var log = zap.L()
	httpSrv := &http.Server{Addr: conf.Addr, Handler: g}
	// 运行 HTTP 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	log.Info("Start to listening the incoming requests on http address", zap.String("addr", conf.Addr))
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()
	return httpSrv
}

// startSecureServer 创建并运行 HTTPS 服务器.
func startSecureServer(g *gin.Engine, conf *conf2.Http) *http.Server {
	// 创建 HTTPS Server 实例
	var log = zap.L()
	httpsSrv := &http.Server{Addr: conf.TslAddr, Handler: g}
	// 运行 HTTPS 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	log.Info("Start to listening the incoming requests on https address", zap.String("addr", conf.TslAddr))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpsSrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err.Error())
			}
		}()
	}
	return httpsSrv
}

// startGRPCServer 创建并运行 GRPC 服务器.
func startGRPCServer(conf *conf2.Config) *grpc.Server {
	var log = zap.L()
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatal("Failed to listen", zap.Error(err))
	}

	// 创建 GRPC Server 实例
	grpcSrv := grpc.NewServer()

	//pb.RegisterMiniBlogServer(grpcsrv, user.New(store.S, nil))

	// 打印一条日志，用来提示 GRPC 服务已经起来，方便排障
	log.Info("Start to listening the incoming requests on grpc address", zap.String("addr", conf.Grpc.Addr))
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
	}()
	return grpcSrv
}
