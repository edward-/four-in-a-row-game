package message

import (
	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
)

type message struct{}

func NewNotifyMessage() Message {
	return &message{}
}

func (c *message) Push(m *entity.Notification) error {
	// TODO push notification
	return nil
}
