package repository

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []string) ([]*entity.UserDTO, error)
	CreateUser(ctx context.Context, userDTO *entity.CreateUserDTO) (string, error)
}
