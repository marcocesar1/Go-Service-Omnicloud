package domain_err

import (
	"errors"
	"fmt"
)

var DuplicatedEmail = errors.New("email already exists")
var InvalidStatus = errors.New("invalid status, valid statuses [IN, OUT]")
var StatusIsTheSame = errors.New("status is the same")

var InvalidPeopleField = errors.New("invalid data")

func InvalidPeopleFieldError(fieldErr string) error {
	return fmt.Errorf("%w: %s", InvalidPeopleField, fieldErr)
}
