package domain_err

import (
	"errors"
	"fmt"
)

var DuplicatedEmail = errors.New("email already exists")
var InvalidStatus = errors.New("invalid status, valid statuses [IN, OUT]")
var StatusIsTheSame = errors.New("status is the same")

var InvalidPeopleField = errors.New("invalid data")

var (
	ErrNameRequired      = errors.New("name field is required")
	ErrNameInvalidLength = errors.New("name field must be between 3 and 80 characters")
	ErrEmailRequired     = errors.New("email field is required")
	ErrEmailInvalid      = errors.New("email field format is invalid")
)

func InvalidPeopleFieldError(err error) error {
	return fmt.Errorf("%w: %w", InvalidPeopleField, err)
}
