package usecase

import (
	"context"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type GameUsecase interface {
	CreateGameExecute(ctx context.Context, game *entity.CreateGameDTO) (string, error)
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
