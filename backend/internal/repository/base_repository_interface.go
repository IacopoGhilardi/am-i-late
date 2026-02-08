package repository

type BaseRepositoryInterface[T any] interface {
	FindAll() ([]T, error)
	Find(id uint) (*T, error)
	Save(entity *T) error
	Delete(id uint) error
}
