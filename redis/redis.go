package redis

import (
	"context"
	"github.com/feitianlove/wechat/config"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	client *redis.Client
)

func NewRedisClient(conf *config.Config) (*redis.Client, error) {
	redisServer := redis.NewClient(&redis.Options{
		Addr:        conf.Redis.Addr,
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    conf.Redis.PoolSize,
		MaxConnAge:  time.Millisecond,
		IdleTimeout: time.Microsecond,
	})
	if _, err := redisServer.Do(context.Background(), "auth", "feitian").Result(); err != nil {
		return nil, err
	}
	client = redisServer
	return redisServer, nil
}
func GetRedisClient() *redis.Client {
	return client
}
