package postgres

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type GameRepository interface {
	GetById(ctx context.Context, gameId string) (*entity.GameDTO, error)
	GetGameActiveByUserId(ctx context.Context, userId string) ([]*entity.GameDTO, error)
	CreateGame(ctx context.Context, gameDTO *entity.CreateGameDTO) (string, error)
	UpdateResult(ctx context.Context, gameId string, result *entity.ResponseTurnDTO) error
}
