package usecases

import "github.com/edward-/four-in-a-row-game/internal/models"

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
}
