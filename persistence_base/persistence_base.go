package persistence_base

import (
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"gorm.io/gorm"
)

// TODO: Add TX support
type Repository[T any] interface {
	GetByID(id uint) (*T, apperr.AppErr)
	GetAll() ([]*T, apperr.AppErr)
	Create(entity *T) apperr.AppErr
	Update(entity *T) apperr.AppErr
	Delete(id uint) apperr.AppErr
}

type gormRepository[T any] struct {
	db                     *gorm.DB
	pem func(err error) apperr.AppErr
}

func NewGormRepository[T any](db *gorm.DB, persistenceErrorMapper func(err error) apperr.AppErr) Repository[T] {
	return &gormRepository[T]{
		db:                     db,
		pem: persistenceErrorMapper,
	}
}

func (g *gormRepository[T]) GetByID(id uint) (*T, apperr.AppErr) {
	var entity T
	err := g.db.First(&entity, id).Error
	return &entity, g.pem(err)
}

func (g *gormRepository[T]) GetAll() ([]*T, apperr.AppErr) {
	var entities []*T
	err := g.db.Find(&entities).Error
	return entities, g.pem(err)
}

func (g *gormRepository[T]) Create(entity *T) apperr.AppErr {
	return g.pem(g.db.Create(entity).Error)
}

func (g *gormRepository[T]) Update(entity *T) apperr.AppErr {
	return g.pem(g.db.Save(entity).Error)
}

func (g *gormRepository[T]) Delete(id uint) apperr.AppErr {
	return g.pem(g.db.Delete(new(T), id).Error)
}
