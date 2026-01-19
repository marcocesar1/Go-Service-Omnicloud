package domain_err

import (
	"errors"
)

var DuplicatedEmail = errors.New("email already exists")
var InvalidStatus = errors.New("invalid status, valid statuses [IN, OUT]")
var StatusIsTheSame = errors.New("status is the same")
