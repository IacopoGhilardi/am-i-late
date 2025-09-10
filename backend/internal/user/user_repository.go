package user

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/db"
)

type UserRepository struct {
	*commons.BaseRepository[User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: commons.NewBaseRepository[User](db.GetDB()),
	}
}
