package app

import (
	"context"
	"fmt"

	"github.com/edward-/four-in-a-row-game/pkg/cache"
	"github.com/edward-/four-in-a-row-game/pkg/config"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	"github.com/edward-/four-in-a-row-game/pkg/database"
	"github.com/edward-/four-in-a-row-game/pkg/logger"
	"github.com/edward-/four-in-a-row-game/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start()
}

type appServer struct {
	ctx    context.Context
	config config.Config
	db     database.Database
	cache  cache.Cache
	logger logger.Logger
}

func NewServer(cfg config.Config, db database.Database, cache cache.Cache, log logger.Logger) Server {
	ctx := context.Background()

	return &appServer{
		ctx:    ctx,
		config: cfg,
		db:     db,
		cache:  cache,
		logger: log,
	}
}

func (s *appServer) Start() {
	ctx := s.ctx
	ctx = contextPkg.SetConfig(ctx, s.config)
	ctx = contextPkg.SetLogger(ctx, s.logger)
	ctx = contextPkg.SetDatabase(ctx, s.db.GetDb())
	ctx = contextPkg.SetCahe(ctx, s.cache.GetCache())

	router := Bootstrap(ctx)

	router.Use(gin.Logger())
	router.Use(middleware.AuthToken())
	router.Use(middleware.SetCorrelationId())
	router.Use(middleware.CommitVersion())
	router.Use(middleware.WithTimeout())

	s.logger.Error(router.Run(fmt.Sprintf(":%d", s.config.App.Port)), "server can not be initialized")
}
