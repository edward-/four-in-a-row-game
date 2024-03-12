package repositories

import "github.com/edward-/four-in-a-row-game/internal/entities"

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}
