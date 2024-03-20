package model

import (
	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
)

type Square struct {
	Column int    `redis:"column"`
	Row    int    `redis:"row"`
	Player string `redis:"player"`
}

type Board struct {
	Squares []*Square `redis:"board"`
}

func (b *Board) ToEntity() *entity.BoardDTO {
	board := new(entity.BoardDTO)

	for _, s := range b.Squares {
		board.Squares[s.Column][s.Row] = s.Player
	}

	return board
}

func BoardToModel(boardDTO *entity.BoardDTO) *Board {
	board := new(Board)
	board.Squares = make([]*Square, 0)

	for column := range vo.BoardColumns {
		for row := range vo.BoardRows {
			player := boardDTO.Squares[column][row]
			s := &Square{Column: column, Row: row, Player: player}
			board.Squares = append(board.Squares, s)
		}
	}

	return board
}
