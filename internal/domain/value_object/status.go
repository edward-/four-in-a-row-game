package valueobject

import "time"

type Status string

const (
	Status_Incompleted Status = "Incompleted"
	Status_Success     Status = "Success"
	Status_InProgress  Status = "InProgress"
)

func GetStatus(isTie bool, winnerId *string, createdAt time.Time) Status {
	success := isCompletedSuccessfully(isTie, winnerId)
	incompletedByTime := isIncompletedByTime(createdAt)

	if !success && incompletedByTime {
		return Status_Incompleted
	}

	if success {
		return Status_Success
	}

	return Status_InProgress
}

func isCompletedSuccessfully(isTie bool, winnerId *string) bool {
	return isTie || winnerId != nil
}

func isIncompletedByTime(createdAt time.Time) bool {
	return createdAt.Before(time.Now().Add(-ActiveGame))
}
