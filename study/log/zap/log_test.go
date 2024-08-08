package zaplog

import (
	"go.uber.org/zap"
	"testing"
)

// https://www.liwenzhou.com/posts/Go/zap/
// https://github.com/uber-go/zap
import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger = CustomerLogger()
	sugarLogger = logger.Sugar()
}
func TestMain(m *testing.M) {
	InitLogger()
	run := m.Run()
	os.Exit(run)
}

func TestStdout(t *testing.T) {
	logger.Error("info", zap.String("name", "lisi"))
}

func TestSugarLogger(t *testing.T) {
	sugarLogger.Infof("info ,name:%s", "lisi")
}
func TestCoreMerge(t *testing.T) {
	encoder := getEncoder()
	// test.log记录全量日志
	logF, _ := os.Create("./test.info.log")
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.InfoLevel)
	// test.err.log记录ERROR级别的日志
	errF, _ := os.Create("./test.err.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zapcore.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)
	logger = zap.New(core, zap.AddCaller())

	logger.Info("info 来了")
	logger.Error("error 来啦")
}

func CustomerLogger() *zap.Logger {
	encoder := getEncoder()
	logWriter := getLogWriter()
	core := zapcore.NewCore(encoder, logWriter, zapcore.InfoLevel)
	//当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，
	//此时日录日志的函数调用链会增加，想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过
	//logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	ws := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(ws)
}
