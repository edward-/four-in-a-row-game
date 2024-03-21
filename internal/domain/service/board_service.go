package service

import (
	"context"
	"errors"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
)

type BoardService interface {
	AnalyzePlay(ctx context.Context, game *entity.GameDTO, board *entity.BoardDTO, turn *entity.TurnDTO) (*entity.ResultTurnDTO, error)
}

type boardService struct{}

func NewBoardService() BoardService {
	return &boardService{}
}

func (b *boardService) AnalyzePlay(ctx context.Context, game *entity.GameDTO, board *entity.BoardDTO, turn *entity.TurnDTO) (result *entity.ResultTurnDTO, err error) {
	result = nextPlayer(game, turn)

	err = addChipToBoard(board, turn)
	if err != nil {
		return nil, err
	}

	if isBoardFull(board) {
		result.Resolution = vo.Resolution_Tie
		result.UserId = vo.SquarEmpty
	}

	if detectFourInRow(board, turn.UserId) {
		result.Resolution = vo.Resolution_Winner
		result.UserId = turn.UserId
	}
	return result, nil
}

func nextPlayer(game *entity.GameDTO, turn *entity.TurnDTO) *entity.ResultTurnDTO {
	userId := turn.UserId
	switch userId {
	case game.UserIds[0]:
		userId = game.UserIds[1]
	case game.UserIds[1]:
		userId = game.UserIds[0]
	}

	return &entity.ResultTurnDTO{
		Resolution: vo.Resolution_Next,
		UserId:     userId,
	}
}

func addChipToBoard(board *entity.BoardDTO, turn *entity.TurnDTO) error {
	if turn.DropItIn > vo.BoardColumns {
		return errors.New("added chip exceeds the maximum allowed")
	}

	for i := range vo.BoardRows {
		player := board.Squares[turn.DropItIn][i]
		if player == vo.SquarEmpty {
			board.Squares[turn.DropItIn][i] = turn.UserId
			return nil
		}
	}

	return errors.New("added chip to column exceeds the maximum allowed")
}

func isBoardFull(board *entity.BoardDTO) bool {
	for i := range vo.BoardRows {
		for j := range vo.BoardColumns {
			if board.Squares[j][i] == "" {
				return false
			}
		}
	}
	return true
}

func detectFourInRow(board *entity.BoardDTO, player string) bool {
	tasks := []taskFunc{
		checkVertical,
		checkHorizontal,
		checkDiagonal1,
		checkDiagonal2,
	}

	return checkDispatcher(tasks, board, player)
}

func checkVertical(board *entity.BoardDTO, player string) bool {
	for i := 0; i <= vo.BoardRows-4; i++ {
		for j := 0; j < vo.BoardColumns; j++ {
			if board.Squares[j][i] == player &&
				board.Squares[j][i+1] == player &&
				board.Squares[j][i+2] == player &&
				board.Squares[j][i+3] == player {
				return true
			}
		}
	}
	return false
}

func checkHorizontal(board *entity.BoardDTO, player string) bool {
	for i := 0; i < vo.BoardRows; i++ {
		for j := 0; j <= vo.BoardColumns-4; j++ {
			if board.Squares[j][i] == player &&
				board.Squares[j+1][i] == player &&
				board.Squares[j+2][i] == player &&
				board.Squares[j+3][i] == player {
				return true
			}
		}
	}
	return false
}

func checkDiagonal1(board *entity.BoardDTO, player string) bool {
	// Check diagonal \
	for i := 0; i <= vo.BoardRows-4; i++ {
		for j := 0; j <= vo.BoardColumns-4; j++ {
			if board.Squares[j][i] == player &&
				board.Squares[j+1][i+1] == player &&
				board.Squares[j+2][i+2] == player &&
				board.Squares[j+3][i+3] == player {
				return true
			}
		}
	}
	return false
}

func checkDiagonal2(board *entity.BoardDTO, player string) bool {
	// Check diagonal /
	for i := 0; i <= vo.BoardRows-4; i++ {
		for j := 3; j < vo.BoardColumns; j++ {
			if board.Squares[j][i] == player &&
				board.Squares[j-1][i+1] == player &&
				board.Squares[j-2][i+2] == player &&
				board.Squares[j-3][i+3] == player {
				return true
			}
		}
	}
	return false
}
