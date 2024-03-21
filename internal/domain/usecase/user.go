package usecase

import (
	"context"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type UserUsecase interface {
	CreateUserExecute(ctx context.Context, user *entity.CreateUserDTO) (string, error)
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
