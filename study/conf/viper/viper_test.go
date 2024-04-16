package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

type Config struct {
	Port    int    `mapstructure:"port1"`
	Version string `mapstructure:"version"`
}

var conf *viper.Viper

func Init() {
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
}

func TestParseConfig(t *testing.T) {
	Init()
	fmt.Println(viper.Get("HTTP.Mode"))
	fmt.Println(viper.Sub("Http"))
}

func TestParseConfig1(t *testing.T) {
	Init()
	var Config = new(Config)
	viper.Unmarshal(Config)
	fmt.Println(Config)
}
