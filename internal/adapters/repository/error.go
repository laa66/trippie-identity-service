package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"gorm.io/gorm"
)
// TODO: more specific error mapper
func MapPostgresError(err error) *apperr.AppErr {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperr.New("record not found")
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Log().Debug("pg err", "code", pgErr.Code)
		switch pgErr.Code {
		case "23505":
			return apperr.New("duplicate record")
		case "23503":
			return apperr.New("foreign key constraint violation")
		case "23502":
			return apperr.New("not null violation")
		default:
			return apperr.New("database error: " + pgErr.Message)
		}
	}

	return apperr.New(err.Error())
}
