package repositories

import "github.com/edward-/four-in-a-row-game/internal/entities"

type UserRepository interface {
	CreateUser(in *entities.CreateUserDTO) error
}
