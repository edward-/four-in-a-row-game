package cache

import "github.com/redis/go-redis/v9"

type Cache interface {
	GetCache() *redis.Client
}
