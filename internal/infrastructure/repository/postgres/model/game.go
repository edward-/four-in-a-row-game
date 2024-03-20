package model

import (
	"time"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
)

type Game struct {
	Id          string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId1     string     `gorm:"column:user_id_1;type:uuid;not null;"`
	UserId2     string     `gorm:"column:user_id_2;type:uuid;not null;"`
	WinnerId    *string    `gorm:"column:winner_id;type:uuid;"`
	IsTie       bool       `gorm:"column:is_tie"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	CompletedAt *time.Time `gorm:"column:completed_at"`
}

func (g *Game) ToEntity() *entity.GameDTO {
	var completedAt int64
	result := new(entity.ResultDTO)

	if g.CompletedAt != nil {
		completedAt = g.CompletedAt.Unix()
	}

	status := vo.GetStatus(g.IsTie, g.WinnerId, g.CreatedAt)
	if status == vo.Status_Success {
		result.Result = vo.GetResult(g.IsTie, g.WinnerId)
		if result.Result == vo.Winner {
			result.Id = g.WinnerId
		}
	}

	userIds := []string{g.UserId1, g.UserId2}

	return &entity.GameDTO{
		Id:          g.Id,
		UserIds:     userIds,
		Status:      status,
		Result:      result,
		CompletedAt: &completedAt,
	}
}

func GameToArrayEntity(games []*Game) []*entity.GameDTO {
	gameDtos := make([]*entity.GameDTO, len(games))
	for i, g := range games {
		gameDtos[i] = g.ToEntity()
	}
	return gameDtos
}
