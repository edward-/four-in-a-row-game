package utils

import (
	"context"
	"fmt"

	"github.com/edward-/four-in-a-row-game/pkg/cache"
	"github.com/edward-/four-in-a-row-game/pkg/config"
	"github.com/edward-/four-in-a-row-game/pkg/database"
)

func DeleteData() {
	cfg := config.GetConfig()
	connection := database.NewPostgresDatabase(&cfg)
	cacheConnection := cache.NewRedisCache(&cfg)

	db := connection.GetDb()
	cache := cacheConnection.GetCache()

	if err := db.Exec("DELETE FROM users; DELETE FROM games;").Error; err != nil {
		panic(fmt.Errorf("failed to delete data. %w", err))
	}

	if err := cache.FlushDB(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("failed to delete cache data. %w", err))
	}
}
