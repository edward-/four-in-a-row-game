package handler

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/usecase"
	httpErrors "github.com/edward-/four-in-a-row-game/internal/interfaces/handler/errors"
	"github.com/edward-/four-in-a-row-game/internal/interfaces/handler/params"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	httpPkg "github.com/edward-/four-in-a-row-game/pkg/http"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ctx          context.Context
	userUsecase  usecase.UserUsecase
	gameUsecase  usecase.GameUsecase
	boardUsecase usecase.BoardUsecase
}

func NewHandler(ctx context.Context,
	userUsecase usecase.UserUsecase,
	gameUsecase usecase.GameUsecase,
	boardUsecase usecase.BoardUsecase,
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
		log.Error(err, httpErrors.ErrInvalidRequestBody)
		response.ResposeBadRequest(httpErrors.ErrInvalidRequestBody)
		return
	}

	userId, err := h.userUsecase.CreateUserExecute(h.ctx, userDTO)
	if err != nil {
		log.Error(err, httpErrors.ErrCreatingUser)
		response.ResposeAbortWithMessage(err, httpErrors.ErrCreatingUser)
		return
	}

	response.ResposeCreatedWithId(userId)
}

func (h *handler) GetUser(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	userId := c.Param(params.UserIdKey)

	user, err := h.userUsecase.GetUserExecute(h.ctx, userId)
	if err != nil {
		log.Error(err, httpErrors.ErrGettingGame)
		response.ResposeAbortWithMessage(err, httpErrors.ErrGettingGame)
		return
	}

	response.ResposeOKWithJSON(user)
}

func (h *handler) CreateGame(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	gameDTO := new(entity.CreateGameDTO)

	if err := c.ShouldBindJSON(gameDTO); err != nil {
		log.Error(err, httpErrors.ErrInvalidRequestBody)
		response.ResposeBadRequest(httpErrors.ErrInvalidRequestBody)
		return
	}

	gameId, err := h.gameUsecase.CreateGameExecute(h.ctx, gameDTO)
	if err != nil {
		log.Error(err, httpErrors.ErrCreatingGame)
		response.ResposeAbortWithMessage(err, httpErrors.ErrCreatingGame)
		return
	}

	response.ResposeCreatedWithId(gameId)
}

func (h *handler) GetGame(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	gameId := c.Param(params.GameIdKey)

	game, err := h.gameUsecase.GetGameExecute(h.ctx, gameId)
	if err != nil {
		log.Error(err, httpErrors.ErrGettingGame)
		response.ResposeAbortWithMessage(err, httpErrors.ErrGettingGame)
		return
	}

	response.ResposeOKWithJSON(game)
}

func (h *handler) GetBoardGame(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	gameId := c.Param(params.GameIdKey)
	board, err := h.boardUsecase.GetBoardExecute(h.ctx, gameId)
	if err != nil {
		log.Error(err, httpErrors.ErrGettingBoard)
		response.ResposeAbortWithMessage(err, httpErrors.ErrGettingBoard)
		return
	}

	response.ResposeOKWithJSON(board)
}

func (h *handler) Turn(c *gin.Context) {
	log := contextPkg.LoggerFromCtx(h.ctx)
	response := httpPkg.NewResponse(c)

	turnDTO := new(entity.TurnDTO)
	if err := c.ShouldBindJSON(turnDTO); err != nil {
		log.Error(err, httpErrors.ErrInvalidRequestBody)
		response.ResposeBadRequest(httpErrors.ErrInvalidRequestBody)
		return
	}

	gameId := c.Param(params.GameIdKey)

	turn, err := h.boardUsecase.TurnExecute(h.ctx, gameId, turnDTO)
	if err != nil {
		log.Error(err, httpErrors.ErrExecutingTurn)
		response.ResposeAbortWithMessage(err, httpErrors.ErrExecutingTurn)
		return
	}

	response.ResposeOKWithJSON(turn)
}
