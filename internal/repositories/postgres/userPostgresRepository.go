package repositories

import (
	"github.com/edward-/four-in-a-row-game/internal/entities"
	"github.com/edward-/four-in-a-row-game/internal/repositories/models"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(userDto *entities.CreateUserDTO) error {
	user := &models.User{
		NickName: userDto.NickName,
		Email:    userDto.Email,
	}

	err := u.db.Create(&user).Error
	if err != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
