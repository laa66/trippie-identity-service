package repository

import (
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepositories struct {
	IdentityRepository *IdentityRepository
}

// TODO: Move dsn to config
func NewPostgresRepositories(dsn string) (*PostgresRepositories, *apperr.AppErr) {
	logger.Log().Debug("creating postgres repositories")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Log().Error("logging", "error", err)
		return nil, apperr.Wrap(err)
	}
	logger.Log().Debug("postgres repositories created")
	return createRepositories(db), nil
}

func createRepositories(db *gorm.DB) *PostgresRepositories {
	return &PostgresRepositories{
		NewIdentityRepository(db),
	}
}
