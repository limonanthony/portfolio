package reviews

import (
	"github.com/limonanthony/portfolio/internal/common"
	"github.com/limonanthony/portfolio/internal/errordefs"
)

func NotFoundErr(id common.Id) error {
	return errordefs.NotFoundf("review with id %d not found", id)
}

func EmailConflictErr(email string) error {
	return errordefs.Conflictf("review with email %s already exists", email)
}
