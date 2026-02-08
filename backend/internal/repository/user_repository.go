package repository

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type UserRepository struct {
	*BaseRepository[model.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[model.User](db.GetDB()),
	}
}

func (r *UserRepository) FindByPublicId(publicId uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.Db.Where("public_id = ?", publicId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) EmailExists(email string) (bool, error) {
	var count int64
	err := r.Db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
