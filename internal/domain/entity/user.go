package entity

type (
	CreateUserDTO struct {
		NickName string `json:"nickName" binding:"required,max=100"`
		Email    string `json:"email" binding:"required,email"`
	}

	UserDTO struct {
		Id        string `json:"id,omitempty"`
		NickName  string `json:"nickName"`
		Email     string `json:"email"`
		CreatedAt string `json:"createdAt"`
	}
)
