package context

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func CacheFromCtx(ctx context.Context) *redis.Client {
	return ctx.Value(cacheKey).(*redis.Client)
}

func SetCahe(ctx context.Context, redisClient *redis.Client) context.Context {
	return context.WithValue(ctx, cacheKey, redisClient)
}
