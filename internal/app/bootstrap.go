package app

import (
	"context"

	usecase2 "github.com/edward-/four-in-a-row-game/internal/domain/usecase"

	"github.com/edward-/four-in-a-row-game/internal/domain/service"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/cache"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/message"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres"
	"github.com/edward-/four-in-a-row-game/internal/interfaces/handler"
	"github.com/gin-gonic/gin"
)

func Bootstrap(ctx context.Context) *gin.Engine {
	router := gin.Default()
	userRepository := postgres.NewUserRepository()
	gameRepository := postgres.NewGameRepository()
	boardRepository := cache.NewBoardRepository()
	notifyMessage := message.NewNotifyMessage()
	boardService := service.NewBoardService()

	userUseCase := usecase2.NewUserUsecase(userRepository, notifyMessage)
	gameUseCase := usecase2.NewGameUsecase(userRepository, gameRepository, boardRepository, notifyMessage)
	boardUseCase := usecase2.NewBoardUsecase(gameRepository, boardRepository, boardService, notifyMessage)

	h := handler.NewHandler(ctx, userUseCase, gameUseCase, boardUseCase)

	router.GET("/ping", h.Ping)
	v1 := router.Group("v1")
	v1.POST("/users", h.CreateUser)
	v1.POST("/users/:userId", h.GetUser)
	v1.POST("/games", h.CreateGame)
	v1.GET("/games/:gameId", h.GetGame)
	v1.GET("/games/:gameId/board", h.GetBoardGame)
	v1.POST("/games/:gameId/turn", h.Turn)

	return router
}
