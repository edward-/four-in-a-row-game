package main

import (
	"github.com/edward-/four-in-a-row-game/internal/app"
	"github.com/edward-/four-in-a-row-game/pkg/cache"
	"github.com/edward-/four-in-a-row-game/pkg/config"
	"github.com/edward-/four-in-a-row-game/pkg/database"
	"github.com/edward-/four-in-a-row-game/pkg/logger"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	cache := cache.NewRedisCache(&cfg)
	log := logger.NewLogger()

	server := app.NewServer(cfg, db, cache, log)
	server.Start()
}
