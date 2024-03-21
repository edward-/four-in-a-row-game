package postgres

import (
	"context"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres/model"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"

	"github.com/pkg/errors"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUsersByIds(ctx context.Context, ids []string) ([]*entity.UserDTO, error) {
	db := contextPkg.DatabaseFromCtx(ctx)
	log := contextPkg.LoggerFromCtx(ctx)

	users := []*model.User{}
	err := db.Where("id IN (?)", ids).Find(&users).Error
	if err != nil {
		msg := "failed getting user"
		log.Error(err, msg)
		return nil, errors.Wrap(err, msg)
	}

	return model.UsersToArrayEntity(users), nil
}

func (u *userRepository) CreateUser(ctx context.Context, userDTO *entity.CreateUserDTO) (string, error) {
	db := contextPkg.DatabaseFromCtx(ctx)
	log := contextPkg.LoggerFromCtx(ctx)

	user := &model.User{
		NickName: userDTO.NickName,
		Email:    userDTO.Email,
	}

	err := db.Where("email = ? and nick_name = ?", user.Email, user.NickName).FirstOrCreate(user).Error
	if err != nil {
		msg := "failed creating user"
		log.Error(err, msg)
		if ok := errors.WithMessage(err, "duplicate key value violates unique constraint"); ok != nil {
			return "", errors.Wrap(err, "email or nickName already exists")
		}
		return "", errors.Wrap(err, msg)
	}
	return user.Id, nil
}
