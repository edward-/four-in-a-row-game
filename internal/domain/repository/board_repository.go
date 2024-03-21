package repository

import (
	"context"
	"time"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type BoardRepository interface {
	Save(ctx context.Context, gameId string, board *entity.BoardDTO, expireTime time.Duration) error
	GetById(ctx context.Context, gameId string) (*entity.BoardDTO, error)
	SetNextTurn(ctx context.Context, gameId string, userId string, expireTime time.Duration) error
	GetNextTurn(ctx context.Context, gameId string) (string, error)
}
