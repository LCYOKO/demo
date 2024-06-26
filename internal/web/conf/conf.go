package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Name         string
	Version      string
	PprofPort    int
	Model        string
	Http         Http
	Grpc         Grpc
	Redis        Redis
	UserDataBase Database
	Log          Log
}

type Http struct {
	Addr    string
	TslAddr string
}

type Grpc struct {
	Addr               string
	ReadTimeoutSecond  int
	WriteTimeoutSecond int
}

type Redis struct {
}

type Database struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	MaxIdeConns int
	MaxConns    int
	MaxIdleTime time.Duration
}

//https://github.com/gin-contrib 第三方插件
type Log struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Level      string
}

func (d *Database) ToMsqlDNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Database)
}

var Conf = new(Config)

func Init() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("conf")
	// 读取配置信息失败
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			// 配置文件被找到，但产生了另外的错误
			return err
		}
	}
	return viper.Unmarshal(Conf)
}
