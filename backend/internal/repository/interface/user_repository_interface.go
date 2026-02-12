package _interface

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type UserRepositoryInterface interface {
	BaseRepositoryInterface[model.User] // Eredita tutti i metodi dell'interfaccia BaseRepository

	FindByPublicId(publicId uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	EmailExists(email string) (bool, error)
}
