package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/validation"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

func (uc *gameUsecase) CreateGameExecute(ctx context.Context, game *entity.CreateGameDTO) (string, error) {
	id := ""

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		users, err := uc.userRepository.GetUsersByIds(ctx, []string{game.UserId1, game.UserId2})
		if err != nil {
			return err
		}
		if err := validation.AreValidUsersToPlay(users); err != nil {
			return err
		}

		id, err = uc.gameRepository.CreateGame(ctx, game)
		if err != nil {
			return err
		}

		board := entity.NewBoard()
		err = uc.boardRepository.Save(ctx, id, board, vo.ActiveGame)
		if err != nil {
			return err
		}

		go uc.notifyMessage.Push(&entity.Notification{Message: "the game has started"})
		return nil
	})
	if err != nil {
		return "", err
	}

	return id, err
}
