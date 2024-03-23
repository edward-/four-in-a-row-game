package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"
)

type UserUsecase interface {
	CreateUserExecute(ctx context.Context, user *entity.CreateUserDTO) (string, error)
	GetUserExecute(ctx context.Context, userId string) (*entity.UserDTO, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	notifyMessage  repository.Message
}

func NewUserUsecase(
	userRepository repository.UserRepository,
	notifyMessage repository.Message,
) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		notifyMessage:  notifyMessage,
	}
}
