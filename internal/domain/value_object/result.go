package valueobject

type Result string

const (
	Tie           Result = "Tie"
	Winner        Result = "Winner"
	UnknownResult Result = "Unknown"
)

func GetResult(isTie bool, winnerId *string) Result {
	if isTie {
		return Tie
	}

	if len(*winnerId) > 0 {
		return Winner
	}

	return UnknownResult
}
