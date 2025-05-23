package repository

import (

	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepositories struct {
	
}

// TODO: Move dsn parameter to config
func NewPostgresRepositories(dsn string) (*PostgresRepositories, *apperr.AppErr) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Log().Error("logging", "error", err)
		return nil, apperr.Wrap(err)
	}
	return createRepository(db), nil
}

func createRepository(db *gorm.DB) *PostgresRepositories {
	return &PostgresRepositories{
		
	}
}