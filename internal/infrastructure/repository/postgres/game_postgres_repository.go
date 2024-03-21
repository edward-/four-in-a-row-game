package postgres

import (
	"context"
	"time"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres/model"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"

	"github.com/pkg/errors"
)

type gameRepository struct {
	table string
}

func NewGameRepository() repository.GameRepository {
	return &gameRepository{
		table: "games",
	}
}

func (u *gameRepository) GetById(ctx context.Context, gameId string) (*entity.GameDTO, error) {
	db := contextPkg.DatabaseFromCtx(ctx)
	log := contextPkg.LoggerFromCtx(ctx)

	game := &model.Game{}

	err := db.Table(u.table).Where("id = ?", gameId).First(&game).Error
	if err != nil {
		msg := "failed getting game"
		log.Error(err, msg)
		return nil, errors.Wrap(err, msg)
	}

	return game.ToEntity(), nil
}

func (u *gameRepository) GetGameActiveByUserId(ctx context.Context, userId string) ([]*entity.GameDTO, error) {
	db := contextPkg.DatabaseFromCtx(ctx)
	log := contextPkg.LoggerFromCtx(ctx)

	games := []*model.Game{}

	query := `select * from games where $1 = ANY(user_ids) and (completed_at is NULL or completed_at = 0)`

	err := db.Table(u.table).Where(query, userId).Find(&games).Error
	if err != nil {
		msg := "failed getting active game by user id"
		log.Error(err, msg)
		return nil, errors.Wrap(err, msg)
	}

	return model.GameToArrayEntity(games), nil
}

func (u *gameRepository) CreateGame(ctx context.Context, gameDTO *entity.CreateGameDTO) (string, error) {
	db := contextPkg.DatabaseFromCtx(ctx)
	log := contextPkg.LoggerFromCtx(ctx)

	game := &model.Game{UserId1: gameDTO.UserId1, UserId2: gameDTO.UserId2}

	err := db.Table(u.table).Create(game).Error
	if err != nil {
		msg := "failed creating game"
		log.Error(err, msg)
		return "", errors.Wrap(err, msg)
	}
	return game.Id, nil
}

func (u *gameRepository) UpdateResult(ctx context.Context, gameId string, result *entity.ResultTurnDTO) error {
	db := contextPkg.DatabaseFromCtx(ctx)
	log := contextPkg.LoggerFromCtx(ctx)

	game := &model.Game{}

	if result.Resolution == vo.Resolution_Winner {
		game.WinnerId = &result.UserId
	}
	if result.Resolution == vo.Resolution_Tie {
		game.IsTie = true
	}
	completedAt := time.Now()
	game.CompletedAt = &completedAt

	err := db.Table(u.table).Where("id = ?", gameId).UpdateColumns(game).Error
	if err != nil {
		msg := "failed updating game"
		log.Error(err, msg)
		return errors.Wrap(err, msg)
	}
	return nil
}
