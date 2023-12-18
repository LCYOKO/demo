package redis

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:            "localhost:6379",
		Password:        "",
		DB:              0,
		MaxRetries:      3,
		MinRetryBackoff: 1,
		PoolSize:        128,
	})
	_, err = RedisClient.Ping().Result()
	return err
}
