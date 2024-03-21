package usecase

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

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
