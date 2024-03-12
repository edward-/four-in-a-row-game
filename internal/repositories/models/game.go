package models

import "time"

type Game struct {
	Id          string
	UserIds     []string
	WinnerId    string
	IsTie       bool
	CreatedAt   bool
	UpdateAt    time.Time
	CompletedAt *time.Time
}
