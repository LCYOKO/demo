package conf

// https://github.com/spf13/viper/blob/master/README.md
// https://liwenzhou.com/posts/Go/viper/
import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

type Config struct {
	Name         string
	Version      string
	PprofPort    int
	Http         Http
	Grpc         Grpc
	Redis        Redis
	UserDataBase Database
}

type Http struct {
	Addr string
}

type Grpc struct {
	Addr               string
	ReadTimeoutSecond  int
	WriteTimeoutSecond int
}

type Redis struct {
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func TestParseConfig1(t *testing.T) {
	// 指定配置文件路径
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../../conf")
	// 读取配置信息失败
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			// 配置文件被找到，但产生了另外的错误
		}
	}
	var conf Config
	viper.Unmarshal(&conf)
	fmt.Println(conf)
}
