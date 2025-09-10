package commons

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Find(id uint) (*T, error) {
	var entity T
	result := r.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	result := r.db.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

// Create or update a entity
func (r *BaseRepository[T]) Save(entity *T) error {
	return r.db.Save(entity).Error
}

// Delete a entity with ID
func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}
