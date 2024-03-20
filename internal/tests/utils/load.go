package utils

import (
	"context"

	"github.com/edward-/four-in-a-row-game/pkg/cache"
	"github.com/edward-/four-in-a-row-game/pkg/config"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	"github.com/edward-/four-in-a-row-game/pkg/database"
	"github.com/edward-/four-in-a-row-game/pkg/logger"
)

func LoadCtx(ctx context.Context) context.Context {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)
	cache := cache.NewRedisCache(&cfg)
	log := logger.NewLogger()

	ctx = contextPkg.SetConfig(ctx, cfg)
	ctx = contextPkg.SetLogger(ctx, log)
	ctx = contextPkg.SetCahe(ctx, cache.GetCache())
	ctx = contextPkg.SetDatabase(ctx, db.GetDb())

	return ctx
}
