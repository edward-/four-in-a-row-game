package entity

import vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"

type (
	CreateGameDTO struct {
		UserId1 string `json:"userId1" binding:"required,uuid4"`
		UserId2 string `json:"userId2" binding:"required,uuid4"`
	}

	GameDTO struct {
		Id          string     `json:"id"`
		UserIds     []string   `json:"user_ids"`
		Status      vo.Status  `json:"status"`
		Result      *ResultDTO `json:"result"`
		CompletedAt *int64     `json:"completedAt,omitempty"`

		Users []*UserDTO
	}

	ResultDTO struct {
		Id     *string   `json:"id"`
		Result vo.Result `json:"result"`
	}
)
