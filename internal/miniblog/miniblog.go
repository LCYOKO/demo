package miniblog

import (
	"context"
	store "demo/internal/miniblog/store/mysql"
	"demo/internal/pkg/known"
	"demo/pkg/auth"
	"demo/pkg/config"
	"demo/pkg/log"
	middleware "demo/pkg/middware"
	"demo/pkg/version/verflag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfgFile string

// NewMiniBlogCommand 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:`,
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()
			log.Init(logOptions())
			defer log.Sync()
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
	cobra.OnInitialize(initConfig)

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
	initConfig()
	if err := store.Init(config.Conf); err != nil {
		return err
	}

	g, err := initGin()
	if err != nil {
		return err
	}
	return runServer(g)
}

func initGin() (g *gin.Engine, err error) {
	auth.Init(viper.GetString("jwt-secret"), known.XUsernameKey)
	gin.SetMode(viper.GetString("runMode"))
	g = gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), middleware.NoCache, middleware.Cors, middleware.Secure, middleware.RequestID()}
	g.Use(mws...)
	if err := installRouters(g); err != nil {
		return nil, err
	}
	return g, nil
}

func runServer(g *gin.Engine) error {
	httpsrv := startInsecureServer(g)
	httpssrv := startSecureServer(g)
	//grpcsrv := startGRPCServer()

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down server ...")

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	if err := httpssrv.Shutdown(ctx); err != nil {
		log.Errorw("Secure Server forced to shutdown", "err", err)
		return err
	}

	//grpcsrv.GracefulStop()
	log.Infow("Server exiting")
	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine) *http.Server {
	httpSrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()
	return httpSrv
}

func startSecureServer(g *gin.Engine) *http.Server {
	httpSsrv := &http.Server{Addr: viper.GetString("tls.addr"), Handler: g}
	log.Infow("Start to listening the incoming requests on https address", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpSsrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalw(err.Error())
			}
		}()
	}

	return httpSsrv
}

//func startGRPCServer() *grpc.Server {
//	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
//	if err != nil {
//		log.Fatalw("Failed to listen", "err", err)
//	}
//	// 创建 GRPC Server 实例
//	grpcSrv := grpc.NewServer()
//	//pb.RegisterMiniBlogServer(grpcSrv, user.New(store.S, nil))
//
//	log.Infow("Start to listening the incoming requests on grpc address", "addr", viper.GetString("grpc.addr"))
//	go func() {
//		if err := grpcSrv.Serve(lis); err != nil {
//			log.Fatalw(err.Error())
//		}
//	}()
//	return grpcSrv
//}
