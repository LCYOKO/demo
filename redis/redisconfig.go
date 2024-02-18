package redis

import (
	"github.com/go-redis/redis"
)

var Client *redis.Client

func Init() (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:            "localhost:6379",
		Password:        "",
		DB:              0,
		MaxRetries:      3,
		MinRetryBackoff: 1,
		PoolSize:        128,
	})
	_, err = Client.Ping().Result()
	return err
}
