package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/validation"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/cache"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/message"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

type GameUsecase interface {
	CreateGameExecute(ctx context.Context, game *entity.CreateGameDTO) (string, error)
}

type gameUsecase struct {
	userRepository  postgres.UserRepository
	gameRepository  postgres.GameRepository
	boardRepository cache.BoardRepository
	notifyMessage   message.Message
}

func NewGameUsecase(
	userRepository postgres.UserRepository,
	gameRepository postgres.GameRepository,
	boardRepository cache.BoardRepository,
	notifyMessage message.Message,
) GameUsecase {
	return &gameUsecase{
		userRepository:  userRepository,
		gameRepository:  gameRepository,
		boardRepository: boardRepository,
		notifyMessage:   notifyMessage,
	}
}

func (u *gameUsecase) CreateGameExecute(ctx context.Context, game *entity.CreateGameDTO) (string, error) {
	id := ""

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		users, err := u.userRepository.GetUsersByIds(ctx, []string{game.UserId1, game.UserId2})
		if err != nil {
			return err
		}
		if err := validation.AreValidUsersToPlay(users); err != nil {
			return err
		}

		id, err = u.gameRepository.CreateGame(ctx, game)
		if err != nil {
			return err
		}

		board := entity.NewBoard()
		err = u.boardRepository.Save(ctx, id, board, vo.ActiveGame)
		if err != nil {
			return err
		}

		go u.notifyMessage.Push(&entity.Notification{Message: "the game has started"})
		return nil
	})
	if err != nil {
		return "", err
	}

	return id, err
}
