package reviews

import "errors"

var (
	ErrEmailConflict = errors.New("email conflict")
	ErrNotFound      = errors.New("review not found")
)
