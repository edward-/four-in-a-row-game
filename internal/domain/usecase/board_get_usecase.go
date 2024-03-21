package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/validation"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

func (uc *boardUsecase) GetBoardExecute(ctx context.Context, gameId string) (*entity.BoardDTO, error) {
	board := new(entity.BoardDTO)

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		game, err := uc.gameRepository.GetById(ctx, gameId)
		if err != nil {
			return err
		}

		err = validation.GameIsAvailableToPlay(game)
		if err != nil {
			return err
		}

		board, err = uc.boardRepository.GetById(ctx, gameId)
		return err
	})
	if err != nil {
		return nil, err
	}

	return board, err
}
