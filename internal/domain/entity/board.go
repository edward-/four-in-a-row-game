package entity

import vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"

type BoardDTO struct {
	Squares [vo.BoardColumns][vo.BoardRows]string `json:"squares"`
}

func NewBoard() *BoardDTO {
	var board [vo.BoardColumns][vo.BoardRows]string

	for column := range vo.BoardColumns {
		for row := range vo.BoardRows {
			board[column][row] = ""
		}
	}

	return &BoardDTO{Squares: board}
}
