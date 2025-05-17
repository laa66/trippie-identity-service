package repository

import (
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/entity"
	persistencebase "github.com/laa66/trippie-identity-service.git/persistence_base"
	"gorm.io/gorm"
)

type IdentityRepository struct {
	*persistencebase.GormRepository[entity.Identity]
}

func NewIdentityRepository(db *gorm.DB) *IdentityRepository {
	ir := persistencebase.NewGormRepository[entity.Identity](db, nil)
	return &IdentityRepository{
		GormRepository: ir,
	}
}