package usecase

import (
	"context"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/service"
)

type BoardUsecase interface {
	TurnExecute(ctx context.Context, gameId string, turn *entity.TurnDTO) (*entity.ResultTurnDTO, error)
	GetBoardExecute(ctx context.Context, gameId string) (*entity.BoardDTO, error)
}

type boardUsecase struct {
	gameRepository  repository.GameRepository
	boardRepository repository.BoardRepository
	boardService    service.BoardService
	notifyMessage   repository.Message
}

func NewBoardUsecase(
	gameRepository repository.GameRepository,
	boardRepository repository.BoardRepository,
	boardService service.BoardService,
	notifyMessage repository.Message,
) BoardUsecase {
	return &boardUsecase{
		gameRepository:  gameRepository,
		boardRepository: boardRepository,
		boardService:    boardService,
		notifyMessage:   notifyMessage,
	}
}
