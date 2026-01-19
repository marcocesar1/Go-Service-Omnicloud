package models

import (
	"slices"
	"time"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PeopleStatus string

const (
	StatusIn  PeopleStatus = "IN"
	StatusOut PeopleStatus = "OUT"
)

type People struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name,omitempty"`
	Email     string        `json:"email" bson:"email,omitempty"`
	Place     string        `json:"place" bson:"place,omitempty"`
	Status    PeopleStatus  `json:"status" bson:"status,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
}

func ValidateStatus(status PeopleStatus) error {
	validStatuses := []PeopleStatus{
		StatusIn,
		StatusOut,
	}
	if !slices.Contains(validStatuses, status) {
		return domain_err.InvalidStatus
	}

	return nil
}
