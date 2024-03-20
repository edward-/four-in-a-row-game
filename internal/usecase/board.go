package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/service"
	"github.com/edward-/four-in-a-row-game/internal/domain/validation"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/cache"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/message"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

type BoardUsecase interface {
	TurnExecute(ctx context.Context, gameId string, turn *entity.TurnDTO) (*entity.ResponseTurnDTO, error)
	GetBoardExecute(ctx context.Context, gameId string) (*entity.BoardDTO, error)
}

type boardUsecase struct {
	gameRepository  postgres.GameRepository
	boardRepository cache.BoardRepository
	boardService    service.BoardService
	notifyMessage   message.Message
}

func NewBoardUsecase(
	gameRepository postgres.GameRepository,
	boardRepository cache.BoardRepository,
	boardService service.BoardService,
	notifyMessage message.Message,
) BoardUsecase {
	return &boardUsecase{
		gameRepository:  gameRepository,
		boardRepository: boardRepository,
		boardService:    boardService,
		notifyMessage:   notifyMessage,
	}
}

func (b *boardUsecase) TurnExecute(ctx context.Context, gameId string, turn *entity.TurnDTO) (*entity.ResponseTurnDTO, error) {
	response := &entity.ResponseTurnDTO{}

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		nextUserId, err := b.boardRepository.GetNextTurn(ctx, gameId)
		if err != nil {
			return err
		}

		err = validation.IsNextTurnCorrectOne(nextUserId, turn.UserId)
		if err != nil {
			return err
		}

		game, err := b.gameRepository.GetById(ctx, gameId)
		if err != nil {
			return err
		}

		err = validation.GameIsAvailableToPlay(game)
		if err != nil {
			return err
		}

		err = validation.UserBelongToGame(game, turn.UserId)
		if err != nil {
			return err
		}

		board, err := b.boardRepository.GetById(ctx, gameId)
		if err != nil {
			return err
		}

		response, err = b.boardService.AnalyzePlay(ctx, game, board, turn)
		if err != nil {
			return err
		}

		if validation.IsNext(response) {
			err = b.boardRepository.SetNextTurn(ctx, gameId, response.UserId, vo.ActiveGame)
			if err != nil {
				return err
			}
		}

		if validation.IsGameOver(response) {
			err = b.gameRepository.UpdateResult(ctx, gameId, response)
			if err != nil {
				return err
			}
			go b.notifyMessage.Push(&entity.Notification{Message: "the game is over"})
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return response, err
}

func (b *boardUsecase) GetBoardExecute(ctx context.Context, gameId string) (*entity.BoardDTO, error) {
	board := new(entity.BoardDTO)

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		game, err := b.gameRepository.GetById(ctx, gameId)
		if err != nil {
			return err
		}

		err = validation.GameIsAvailableToPlay(game)
		if err != nil {
			return err
		}

		board, err = b.boardRepository.GetById(ctx, gameId)
		return err
	})
	if err != nil {
		return nil, err
	}

	return board, err
}
