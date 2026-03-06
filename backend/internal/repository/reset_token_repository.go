package repository

import (
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type ResetTokenRepository struct {
	*BaseRepository[model.PasswordResetToken]
}

func NewResetTokenRepository() *ResetTokenRepository {
	return &ResetTokenRepository{
		BaseRepository: NewBaseRepository[model.PasswordResetToken](db.GetDB()),
	}
}

func (r *ResetTokenRepository) FindByToken(token string) (*model.PasswordResetToken, error) {
	var resetToken model.PasswordResetToken
	err := r.Db.Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}
