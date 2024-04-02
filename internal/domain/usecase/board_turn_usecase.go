package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/validation"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

func (uc *boardUsecase) TurnExecute(ctx context.Context, gameId string, turn *entity.TurnDTO) (*entity.ResultTurnDTO, error) {
	response := &entity.ResultTurnDTO{}

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		err = uc.getAndValidateCurrentTurn(ctx, gameId, turn)
		if err != nil {
			return err
		}

		game, err := uc.getAndValidateGame(ctx, gameId, turn.UserId)
		if err != nil {
			return err
		}

		response, err = uc.processGame(ctx, game, turn)
		return err
	})
	if err != nil {
		return nil, err
	}

	return response, err
}

func (uc *boardUsecase) getAndValidateCurrentTurn(ctx context.Context, gameId string, turn *entity.TurnDTO) error {
	nextUserId, err := uc.boardRepository.GetTurn(ctx, gameId)
	if err != nil {
		return err
	}

	err = validation.IsNextTurnCorrectOne(nextUserId, turn.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (uc *boardUsecase) getAndValidateGame(ctx context.Context, gameId string, userId string) (*entity.GameDTO, error) {
	game, err := uc.gameRepository.GetById(ctx, gameId)
	if err != nil {
		return nil, err
	}

	err = validation.GameIsAvailableToPlay(game)
	if err != nil {
		return nil, err
	}

	err = validation.UserBelongToGame(game, userId)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (uc *boardUsecase) processGame(ctx context.Context, game *entity.GameDTO, turn *entity.TurnDTO) (*entity.ResultTurnDTO, error) {
	board, err := uc.boardRepository.GetById(ctx, game.Id)
	if err != nil {
		return nil, err
	}

	result, err := uc.boardService.AnalyzePlay(ctx, game, board, turn)
	if err != nil {
		return nil, err
	}

	err = uc.boardRepository.Save(ctx, game.Id, board, vo.ActiveGame)
	if err != nil {
		return nil, err
	}

	if validation.IsNext(result) {
		err = uc.boardRepository.SetNextTurn(ctx, game.Id, result.UserId, vo.ActiveGame)
		if err != nil {
			return nil, err
		}
	}

	if validation.IsGameOver(result) {
		err = uc.gameRepository.UpdateResult(ctx, game.Id, result)
		if err != nil {
			return nil, err
		}
		go uc.notifyMessage.Push(&entity.Notification{Message: "the game is over"})
	}

	return result, nil
}
