package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Name         string `mapstructure:"appName"`
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

var once sync.Once
var Conf = &Config{}

func Load() {
	once.Do(func() {
		viper.SetConfigFile("./conf/config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
		// 监控配置文件变化
		viper.WatchConfig()
		// 注意！！！配置文件发生变化后要同步到全局变量Conf
		viper.OnConfigChange(func(in fsnotify.Event) {
			if err := viper.Unmarshal(Conf); err != nil {
				panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
			}
		})
	})
}
