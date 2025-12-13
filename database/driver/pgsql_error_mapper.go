package database

import (
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type DbErrorCode string

const (
	// Common
	ErrUnknown DbErrorCode = "UNKNOWN_ERROR"

	// DB specific domain errors
	ErrDBUnique       DbErrorCode = "UNIQUE_ERROR"
	ErrDBForeignKey   DbErrorCode = "FOREIGN_KEY_ERROR"
	ErrDBNotNull      DbErrorCode = "NOT_NULL_ERROR"
	ErrDBCheck        DbErrorCode = "CHECK_ERROR"
	ErrDBInvalidInput DbErrorCode = "INVALID_INPUT_ERROR"
)

type DatabaseError struct {
	Code    DbErrorCode
	Message string
	Field   string // Optional: which column caused the error
	Details string // Optional: raw details
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func MapPgError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return err // not a postgres error
	}

	switch pgErr.Code {

	case pgerrcode.UniqueViolation:
		return &DatabaseError{
			Code:    ErrDBUnique,
			Field:   pgErr.ConstraintName,
			Message: "duplicate value violates unique constraint",
			Details: pgErr.Detail,
		}

	case pgerrcode.ForeignKeyViolation:
		return &DatabaseError{
			Code:    ErrDBForeignKey,
			Field:   pgErr.ConstraintName,
			Message: "foreign key constraint failed",
			Details: pgErr.Detail,
		}

	case pgerrcode.NotNullViolation:
		return &DatabaseError{
			Code:    ErrDBNotNull,
			Field:   pgErr.ColumnName,
			Message: "required field is missing",
			Details: pgErr.Detail,
		}

	case pgerrcode.CheckViolation:
		return &DatabaseError{
			Code:    ErrDBCheck,
			Field:   pgErr.ConstraintName,
			Message: "check constraint validation failed",
			Details: pgErr.Detail,
		}

	case pgerrcode.InvalidTextRepresentation:
		return &DatabaseError{
			Code:    ErrDBInvalidInput,
			Field:   pgErr.ColumnName,
			Message: "invalid input format",
			Details: pgErr.Detail,
		}

	default:
		return &DatabaseError{
			Code:    ErrUnknown,
			Message: pgErr.Message,
			Details: pgErr.Detail,
		}
	}
}
