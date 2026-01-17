package domain_err

import err "errors"

var ErrNotFound = err.New("document not found")
var ErrDuplicatedDoc = err.New("duplicated document")
var ErrInvalidObjectId = err.New("invalid object id")
