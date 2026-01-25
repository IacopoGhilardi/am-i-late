package repository

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type UserRepository struct {
	*commons.BaseRepository[model.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: commons.NewBaseRepository[model.User](db.GetDB()),
	}
}
