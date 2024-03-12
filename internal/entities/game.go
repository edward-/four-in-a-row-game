package entities

type GameDto struct {
	Id        string            `json:"id"`
	Users     map[string]string `json:"users"`
	Status    string            `json:"status"`
	CreatedAt int64             `json:"createdAt"`
	UpdateAt  int64             `json:"updatedAt"`
	Completed *int64            `json:"completedAt,omitempty"`
}
