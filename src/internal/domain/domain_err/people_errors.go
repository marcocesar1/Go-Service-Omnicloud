package domain_err

import "errors"

var DuplicatedEmail = errors.New("email already exists")
