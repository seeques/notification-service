package storage

import (
	"time"
)

type Template struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Preference struct {
	ID        int       `json:"id"`
	UserID    string    `json:"userID"`
	Channel   string    `json:"channel"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Log struct {
	ID        int       `json:"id"`
	JobID     int       `json:"jobID"`
	Channel   string    `json:"channel"`
	Recipient string    `json:"recipient"`
	Status    string    `json:"status"`
	Attempts  int       `json:"attempts"`
	ErrMsg    string    `json:"errorMessage"`
	CreatedAt time.Time `json:"createdAt"`
	SentAt    time.Time `json:"sentAt"`
}
