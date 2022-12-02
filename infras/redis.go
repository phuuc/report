package infras

import (
	"context"
	"log"
	"time"

	"github.com/finnpn/overview/config"
	"github.com/go-redis/redis/v9"
)

func InitRedis(redisCfg *config.RedisConfig) (*redis.Client, error) {
	var redisClient *redis.Client

	opts, err := redis.ParseURL(redisCfg.ConnectionURL)
	if err != nil {
		return nil, err
	}

	opts.PoolSize = redisCfg.PoolSize
	opts.DialTimeout = time.Duration(redisCfg.DialTimeoutSeconds) * time.Second
	opts.ReadTimeout = time.Duration(redisCfg.ReadTimeoutSeconds) * time.Second
	opts.WriteTimeout = time.Duration(redisCfg.WriteTimeoutSeconds) * time.Second

	redisClient = redis.NewClient(opts)

	cmd := redisClient.Ping(context.Background())
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	log.Println("connect to redis successful")
	return redisClient, nil
}
