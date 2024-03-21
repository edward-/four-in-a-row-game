package entity

import vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"

type TurnDTO struct {
	UserId   string `json:"userId" binding:"required,uuid4"`
	DropItIn int    `json:"dropItIn"`
}

type ResultTurnDTO struct {
	Resolution vo.Resolution `json:"resolution"`
	UserId     string        `json:"user_id"`
}
