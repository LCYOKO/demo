package zaplog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"testing"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger = CustomerLogger()
	sugarLogger = logger.Sugar()
}

func TestStdout(t *testing.T) {
	InitLogger()
	logger.Info("info", zap.String("name", "lisi"))
}

func TestSugarLogger(t *testing.T) {
	InitLogger()
	sugarLogger.Infof("info ,name:%s", "lisi")
}

func CustomerLogger() *zap.Logger {
	encoder := getEncoder()
	logWriter := getLogWriter()
	core := zapcore.NewCore(encoder, logWriter, zapcore.InfoLevel)
	return zap.New(core)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime =zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
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
