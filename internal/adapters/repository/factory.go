package repository

import (
	"fmt"

	"github.com/laa66/trippie-identity-service.git/config"
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"github.com/laa66/trippie-identity-service.git/internal/adapters/logger"
	"github.com/laa66/trippie-identity-service.git/internal/core/ports/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepositories interface {
	GetIdentityRepository() repository.IdentityRepository
}

type postgresRepositories struct {
	IdentityRepository repository.IdentityRepository
}

func NewPostgresRepositories() (*postgresRepositories, *apperr.AppErr) {
	logger.Log().Debug("creating postgres repositories")

	db, err := gorm.Open(postgres.Open(createDSN()))
	if err != nil {
		logger.Log().Error("logging", "error", err)
		return nil, apperr.Wrap(err)
	}
	logger.Log().Debug("postgres repositories created")
	return createRepositories(db), nil
}

func (p *postgresRepositories) GetIdentityRepository() repository.IdentityRepository {
	return p.IdentityRepository
}

func createRepositories(db *gorm.DB) *postgresRepositories {
	return &postgresRepositories{
		NewIdentityRepository(db),
	}
}

func createDSN() string {
	cfg := config.GetConfig().DB
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
}
