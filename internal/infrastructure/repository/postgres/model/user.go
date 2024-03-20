package model

import "github.com/edward-/four-in-a-row-game/internal/domain/entity"

type User struct {
	Id       string `gorm:"type:uuid;default:uuid_generate_v4()"`
	NickName string `gorm:"column:nick_name"`
	Email    string `gorm:"column:email"`
}

func (u *User) ToEntity() *entity.UserDTO {
	return &entity.UserDTO{
		Id:       u.Id,
		NickName: u.NickName,
		Email:    u.Email,
	}
}

func UsersToArrayEntity(users []*User) []*entity.UserDTO {
	userDtos := make([]*entity.UserDTO, len(users))
	for i, u := range users {
		userDtos[i] = u.ToEntity()
	}
	return userDtos
}
