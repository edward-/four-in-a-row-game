package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/edward-/four-in-a-row-game/pkg/config"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	Client *redis.Client
}

func NewRedisCache(cfg *config.Config) Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port),
		Password:     cfg.Cache.Password,
		DB:           cfg.Cache.Db,
		Protocol:     cfg.Cache.Protocol,
		MinIdleConns: cfg.Cache.MinIdleConns,
		PoolSize:     cfg.Cache.PoolSize,
		PoolTimeout:  time.Duration(cfg.Cache.PoolTimeout),
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &redisCache{Client: redisClient}
}

func (r *redisCache) GetCache() *redis.Client {
	return r.Client
}
