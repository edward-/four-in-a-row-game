package repository

import (
	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type Message interface {
	Push(m *entity.Notification) error
}
