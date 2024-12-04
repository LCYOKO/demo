package config

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

type Redis struct {
}

var _defaultPath = ".././conf/"
var _defaultConf = &Config{}
var once sync.Once
var Conf = _defaultConf

func Init(path string) {
	once.Do(func() {
		viper.AddConfigPath(_defaultPath)
		if len(path) > 0 {
			viper.AddConfigPath(path)
		}
		viper.SetConfigFile("./conf/config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("read conf file failed error: %v \n", err))
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
