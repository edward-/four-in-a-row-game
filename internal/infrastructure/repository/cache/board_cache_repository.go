package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/cache/model"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	"github.com/pkg/errors"
)

type boardRepository struct{}

func NewBoardRepository() repository.BoardRepository {
	return &boardRepository{}
}

func (u *boardRepository) GetById(ctx context.Context, gameId string) (*entity.BoardDTO, error) {
	cache := contextPkg.CacheFromCtx(ctx)

	boardBytes, err := cache.Get(ctx, gameId).Bytes()
	if err != nil {
		msg := "failed getting board from cache"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return nil, errors.Wrap(err, msg)
	}

	board := &model.Board{}
	if err = json.Unmarshal(boardBytes, &board); err != nil {
		msg := "failed recovering board from cache"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return nil, errors.Wrap(err, msg)
	}

	return board.ToEntity(), nil
}

func (u *boardRepository) Save(ctx context.Context, gameId string, boardDTO *entity.BoardDTO, expireTime time.Duration) error {
	cache := contextPkg.CacheFromCtx(ctx)

	board := model.BoardToModel(boardDTO)

	boardBytes, err := json.Marshal(&board)
	if err != nil {
		msg := "failed board converting to bytes"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return errors.Wrap(err, msg)
	}

	err = cache.Set(ctx, gameId, boardBytes, expireTime).Err()
	if err != nil {
		msg := "failed saving board in cache"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return errors.Wrap(err, msg)
	}

	return nil
}

func (u *boardRepository) SetNextTurn(ctx context.Context, gameId string, userId string, expireTime time.Duration) error {
	cache := contextPkg.CacheFromCtx(ctx)

	userIdBytes, err := json.Marshal(&userId)
	if err != nil {
		msg := "failed` converting to bytes"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return errors.Wrap(err, msg)
	}

	err = cache.Set(ctx, gameId+"_next", userIdBytes, expireTime).Err()
	if err != nil {
		msg := "failed saving next to cache"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return errors.Wrap(err, msg)
	}

	return nil
}

func (u *boardRepository) GetNextTurn(ctx context.Context, gameId string) (string, error) {
	cache := contextPkg.CacheFromCtx(ctx)

	bytes, err := cache.Get(ctx, gameId+"_next").Bytes()
	if err != nil {
		if err.Error() == "redis: nil" {
			return "", nil
		}
		msg := "failed getting next from cache"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return "", errors.Wrap(err, msg)
	}

	var userId string
	if err = json.Unmarshal(bytes, &userId); err != nil {
		msg := "failed recovering next from cache"
		contextPkg.LoggerFromCtx(ctx).Error(err, msg)
		return "", errors.Wrap(err, msg)
	}

	return userId, nil
}
