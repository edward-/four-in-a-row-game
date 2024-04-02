package usecase

import (
	"context"
	"errors"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/pkg/transaction"
)

func (u *userUsecase) GetUserExecute(ctx context.Context, userId string) (user *entity.UserDTO, err error) {
	err = transaction.NewTransaction(ctx, func(ctx context.Context) (err error) {
		users, err := u.userRepository.GetUsersByIds(ctx, []string{userId})
		if err != nil {
			return err
		}
		if len(users) == 0 {
			return errors.New("user not found")
		}

		user = users[0]
		return nil
	})

	return user, err
}
