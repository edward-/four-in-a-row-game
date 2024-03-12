package entities

type (
	CreateUserDTO struct {
		NickName string `json:"nickName" validate:"required"`
		Email    string `json:"number" validate:"required,email"`
	}

	UserDto struct {
		Id        string `json:"id"`
		NickName  string `json:"nickName"`
		Email     string `json:"number"`
		CreatedAt string `json:"createdAt"`
	}
)
