package models

import (
	"time"
)

type PeopleStatus string

const (
	StatusIn  PeopleStatus = "IN"
	StatusOut PeopleStatus = "OUT"
)

type People struct {
	ID        string       `json:"id" `
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Place     string       `json:"place"`
	Status    PeopleStatus `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
