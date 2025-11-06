package dbstore

import (
	"database/sql"
	"errors"

	"github.com/Ilja-R/library-service/internal/errs"
)

func (p *BookStorage) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}