package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"
)

type GameUsecase interface {
	CreateGameExecute(ctx context.Context, game *entity.CreateGameDTO) (string, error)
	GetGameExecute(ctx context.Context, gameId string) (*entity.GameDTO, error)
}

type gameUsecase struct {
	userRepository  repository.UserRepository
	gameRepository  repository.GameRepository
	boardRepository repository.BoardRepository
	notifyMessage   repository.Message
}

func NewGameUsecase(
	userRepository repository.UserRepository,
	gameRepository repository.GameRepository,
	boardRepository repository.BoardRepository,
	notifyMessage repository.Message,
) GameUsecase {
	return &gameUsecase{
		userRepository:  userRepository,
		gameRepository:  gameRepository,
		boardRepository: boardRepository,
		notifyMessage:   notifyMessage,
	}
}
