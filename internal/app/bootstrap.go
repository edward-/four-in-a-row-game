package app

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/service"
	"github.com/edward-/four-in-a-row-game/internal/handler"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/cache"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/message"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres"
	"github.com/edward-/four-in-a-row-game/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Bootstrap(ctx context.Context) *gin.Engine {
	router := gin.Default()
	userRepository := postgres.NewUserRepository()
	gameRepository := postgres.NewGameRepository()
	boardRepository := cache.NewBoardRepository()
	notifyMessage := message.NewNotifyMessage()
	boardService := service.NewBoardService()

	userUseCase := usecase.NewUserUsecase(userRepository, notifyMessage)
	gameUseCase := usecase.NewGameUsecase(userRepository, gameRepository, boardRepository, notifyMessage)
	boardUseCase := usecase.NewBoardUsecase(gameRepository, boardRepository, boardService, notifyMessage)

	h := handler.NewHandler(ctx, userUseCase, gameUseCase, boardUseCase)

	v1 := router.Group("v1")
	v1.POST("/users", h.CreateUser)
	v1.POST("/games", h.CreateGame)
	v1.GET("/games/:gameId/board", h.GetBoardGame)
	v1.POST("/games/:gameId/turn", h.Turn)

	return router
}
