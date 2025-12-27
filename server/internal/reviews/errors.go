package reviews

import "github.com/limonanthony/portfolio/internal/errordefs"

var (
	ErrEmailConflict = errordefs.NewConflict("email already exists", nil)
	ErrNotFound      = errordefs.NewNotFound("review not found", nil)
)
