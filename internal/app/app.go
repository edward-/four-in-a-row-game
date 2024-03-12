package app

import (
	"fmt"

	"github.com/edward-/four-in-a-row-game/internal/handlers"
	"github.com/edward-/four-in-a-row-game/internal/repositories"
	"github.com/edward-/four-in-a-row-game/internal/usecases"
	"github.com/edward-/four-in-a-row-game/pkg/config"
	"github.com/edward-/four-in-a-row-game/pkg/logger"
	"github.com/edward-/four-in-a-row-game/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server interface {
	Start()
}

type appServer struct {
	app    *gin.Engine
	db     *gorm.DB
	cfg    *config.Config
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *gorm.DB) Server {
	return &appServer{
		app:    gin.New(),
		db:     db,
		cfg:    cfg,
		logger: logger.NewLogger(),
	}
}

func (s *appServer) Start() {
	s.initializeHandler()
	s.app.Use(gin.Logger())
	s.app.Use(middleware.SetCorrelationUUIDMiddleware())

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)

	s.logger.Error(s.app.Run(serverUrl).Error())
}

func (s *appServer) initializeHandler() {
	handler := middleware.SetCorrelationUUIDMiddleware(s.app)

	// Initialize all layers
	cockroachPostgresRepository := repositories.NewUserRepository(s.db)
	cockroachFCMMessaging := repositories.NewCockroachFCMMessaging()

	cockroachUsecase := usecases.NewCockroachUsecaseImpl(
		cockroachPostgresRepository,
		cockroachFCMMessaging,
	)

	cockroachHttpHandler := handlers.NewCockroachHttpHandler(cockroachUsecase)

	// Routers
	cockroachRouters := s.app.Group("v1/cockroach")
	cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
}
