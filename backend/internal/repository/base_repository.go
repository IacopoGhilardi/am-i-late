package repository

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	Db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{Db: db}
}

func (r *BaseRepository[T]) Find(id uint) (*T, error) {
	var entity T
	result := r.Db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	result := r.Db.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (r *BaseRepository[T]) Save(entity *T) error {
	return r.Db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	return r.Db.Delete(&entity, id).Error
}
