package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

func (uc *gameUsecase) GetGameExecute(ctx context.Context, gameId string) (game *entity.GameDTO, err error) {
	err = transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		game, err = uc.gameRepository.GetById(ctx, gameId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return game, err
	}

	return game, err
}
