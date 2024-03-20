package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/message"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

type UserUsecase interface {
	CreateUserExecute(ctx context.Context, user *entity.CreateUserDTO) (string, error)
}

type userUsecase struct {
	userRepository postgres.UserRepository
	notifyMessage  message.Message
}

func NewUserUsecase(
	userRepository postgres.UserRepository,
	notifyMessage message.Message,
) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		notifyMessage:  notifyMessage,
	}
}

func (u *userUsecase) CreateUserExecute(ctx context.Context, user *entity.CreateUserDTO) (string, error) {
	id := ""

	err := transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		id, err = u.userRepository.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		go u.notifyMessage.Push(&entity.Notification{Message: "user created"})
		return nil
	})

	return id, err
}
