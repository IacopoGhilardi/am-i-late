package _interface

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type UserServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByPublicId(publicId uuid.UUID) (*model.User, error)
	CreateUser(u *model.User) (*model.User, error)
	DeleteUser(id uint) error
}
