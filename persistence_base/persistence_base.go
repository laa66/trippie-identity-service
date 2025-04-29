package persistence_base

import (
	apperr "github.com/laa66/trippie-identity-service.git/error"
	"gorm.io/gorm"
)

// TODO: Move to lib, add TX support, create persistence/postgre_persistence that initializes db connection and build other persistences
type Repository[T any] interface {
	GetByID(id uint) (*T, *apperr.AppErr)
	GetAll() ([]*T, *apperr.AppErr)
	Create(entity *T) *apperr.AppErr
	Update(entity *T) *apperr.AppErr
	Delete(id uint) *apperr.AppErr
}

type gormRepository[T any] struct {
	db   *gorm.DB
	pem  func(err error) *apperr.AppErr
	isTx bool
}

func NewGormRepository[T any](db *gorm.DB, persistenceErrorMapper func(err error) *apperr.AppErr) Repository[T] {
	db.Begin()
	return &gormRepository[T]{
		db:  db,
		pem: persistenceErrorMapper,
	}
}

func (g *gormRepository[T]) Begin() (Repository[T], *apperr.AppErr) {
	tx := g.db.Begin()
	if tx.Error != nil {
		return nil, g.pem(tx.Error)
	}
	return &gormRepository[T]{db: tx, pem: g.pem, isTx: true}, nil
}

func (g *gormRepository[T]) Commit() *apperr.AppErr {
	if g.isTx {
		return g.pem(g.db.Commit().Error)
	}
	return nil
}

func (g *gormRepository[T]) Rollback() *apperr.AppErr {
	if g.isTx {
		return g.pem(g.db.Rollback().Error)
	}
	return nil
}

func (g *gormRepository[T]) GetByID(id uint) (*T, *apperr.AppErr) {
	var entity T
	err := g.db.First(&entity, id).Error
	return &entity, g.pem(err)
}

func (g *gormRepository[T]) GetAll() ([]*T, *apperr.AppErr) {
	var entities []*T
	err := g.db.Find(&entities).Error
	return entities, g.pem(err)
}

func (g *gormRepository[T]) Create(entity *T) *apperr.AppErr {
	return g.pem(g.db.Create(entity).Error)
}

func (g *gormRepository[T]) Update(entity *T) *apperr.AppErr {
	return g.pem(g.db.Save(entity).Error)
}

func (g *gormRepository[T]) Delete(id uint) *apperr.AppErr {
	return g.pem(g.db.Delete(new(T), id).Error)
}
