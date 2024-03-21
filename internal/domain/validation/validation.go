package validation

import (
	"errors"
	"slices"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
)

func GameIsAvailableToPlay(game *entity.GameDTO) error {
	switch game.Status {
	case vo.Status_InProgress:
		return nil
	case vo.Status_Success:
		return errors.New("game is finished")
	case vo.Status_Incompleted:
		return errors.New("game is incompleted and can't be played")
	default:
		return errors.New("unknown status")
	}
}

func IsGameOver(response *entity.ResultTurnDTO) bool {
	return response.Resolution == vo.Resolution_Winner || response.Resolution == vo.Resolution_Tie
}

func IsNext(response *entity.ResultTurnDTO) bool {
	return response.Resolution == vo.Resolution_Next
}

func AreValidUsersToPlay(users []*entity.UserDTO) error {
	if len(users) == 2 {
		return nil
	}
	return errors.New("invalid users")
}

func UserBelongToGame(game *entity.GameDTO, userId string) error {
	if slices.Contains(game.UserIds, userId) {
		return nil
	}
	return errors.New("user don't belong to game")
}

func IsNextTurnCorrectOne(nextUserId, userId string) error {
	if nextUserId == "" {
		return nil
	}
	if nextUserId == userId {
		return nil
	}
	return errors.New("invalid next turn")
}
