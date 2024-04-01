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
	viper.SetConfigFile("config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func TestParseConfig(t *testing.T) {
	Init()
	fmt.Println(viper.Get("version"))
}

func TestParseConfig1(t *testing.T) {
	Init()
	var Config = new(Config)
	viper.Unmarshal(Config)
	fmt.Println(Config)
}
