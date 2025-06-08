package repository

import (
	"github.com/laa66/trippie-identity-service.git/internal/core/domain/entity"
	persistencebase "github.com/laa66/trippie-identity-service.git/persistence_base"
)

type IdentityRepository interface {
	persistencebase.Repository[entity.Identity]
}