package handler

import (
	"context"
	usecase2 "github.com/edward-/four-in-a-row-game/internal/domain/usecase"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	httpPkg "github.com/edward-/four-in-a-row-game/pkg/http"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ctx          context.Context
	userUsecase  usecase2.UserUsecase
	gameUsecase  usecase2.GameUsecase
	boardUsecase usecase2.BoardUsecase
}

func NewHandler(ctx context.Context,
	userUsecase usecase2.UserUsecase,
	gameUsecase usecase2.GameUsecase,
	boardUsecase usecase2.BoardUsecase,
) Handler {
	return &handler{
		ctx:          ctx,
		userUsecase:  userUsecase,
		gameUsecase:  gameUsecase,
		boardUsecase: boardUsecase,
	}
}

func (h *handler) Ping(c *gin.Context) {
	response := httpPkg.NewResponse(c)
	response.ResposeOKWithJSON(gin.H{"response": "pong"})
}

func (h *handler) CreateUser(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	userDTO := new(entity.CreateUserDTO)

	if err := c.ShouldBindJSON(userDTO); err != nil {
		log.Error(err, "body invalid")
		response.ResposeBadRequest("body invalid")
		return
	}

	userId, err := h.userUsecase.CreateUserExecute(h.ctx, userDTO)
	if err != nil {
		log.Error(err, "could not create the user")
		response.ResposeAbortWithMessage(err, "could not create the user")
		return
	}

	response.ResposeCreatedWithId(userId)
}

func (h *handler) CreateGame(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	gameDTO := new(entity.CreateGameDTO)

	if err := c.ShouldBindJSON(gameDTO); err != nil {
		log.Error(err, "body invalid")
		response.ResposeBadRequest("body invalid")
		return
	}

	gameId, err := h.gameUsecase.CreateGameExecute(h.ctx, gameDTO)
	if err != nil {
		log.Error(err, "could not create the game")
		response.ResposeAbortWithMessage(err, "could not create the game")
		return
	}

	response.ResposeCreatedWithId(gameId)
}

func (h *handler) GetBoardGame(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	gameId := c.Param("gameId")
	board, err := h.boardUsecase.GetBoardExecute(h.ctx, gameId)
	if err != nil {
		log.Error(err, "could not get the board")
		response.ResposeAbortWithMessage(err, "could not get the board")
		return
	}

	response.ResposeOKWithJSON(board)
}

func (h *handler) Turn(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	turnDTO := new(entity.TurnDTO)
	if err := c.ShouldBindJSON(turnDTO); err != nil {
		log.Error(err, "body invalid")
		response.ResposeBadRequest("body invalid")
		return
	}

	gameId := c.Param("gameId")

	turn, err := h.boardUsecase.TurnExecute(h.ctx, gameId, turnDTO)
	if err != nil {
		log.Error(err, "could not do next move")
		response.ResposeAbortWithMessage(err, "could not do next move")
		return
	}

	response.ResposeOKWithJSON(turn)
}
