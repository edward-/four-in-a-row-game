package message

import (
	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	"github.com/edward-/four-in-a-row-game/internal/domain/repository"
)

type message struct{}

func NewNotifyMessage() repository.Message {
	return &message{}
}

func (c *message) Push(m *entity.Notification) error {
	// TODO push notification
	return nil
}
