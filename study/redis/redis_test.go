package redis

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var Cli *redis.Client
var SentinelClient *redis.Client
var ClusterClient *redis.ClusterClient

func initNormal() {
	Cli = redis.NewClient(&redis.Options{
		Addr:     "114.55.147.178:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
}

func initSentinel() {
	SentinelClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
}

func initCluster() {
	ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

		// 若要根据延迟或随机路由命令，请启用以下命令之一
		// RouteByLatency: true,
		// RouteRandomly: true,
	})
}

func TestNormal(t *testing.T) {
	initNormal()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := Cli.Get(ctx, "key").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := Cli.Get(ctx, "key")
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = Cli.Set(ctx, "key", 10, time.Hour).Err()

	// 直接执行命令获取值
	value := Cli.Get(ctx, "key").Val()
	fmt.Println(value)
}

func TestDo(t *testing.T) {
	initNormal()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接执行命令获取错误
	err := Cli.Do(ctx, "set", "key", 10, "EX", 3600).Err()
	fmt.Println(err)

	// 执行命令获取结果
	val, err := Cli.Do(ctx, "get", "key").Result()
	fmt.Println(val, err)
}

func TestNilError(t *testing.T) {
	initNormal()
	fromRedis, err := getValueFromRedis("test", "wu")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(fromRedis)
}

// getValueFromRedis redis.Nil判断
func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := Cli.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		// 出其他错了
		return "", err
	}
	return val, nil
}
