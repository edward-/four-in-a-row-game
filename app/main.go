package app

import (
	"github.com/edward-/four-in-a-row-game/internal/app"
	"github.com/edward-/four-in-a-row-game/pkg/config"
	"github.com/edward-/four-in-a-row-game/pkg/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	app.NewServer(&cfg, db.GetDb()).Start()
}
