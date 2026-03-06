package _interface

import "github.com/iacopoGhilardi/amILate/internal/model"

type PasswordResetTokenRepositoryInterface interface {
	BaseRepositoryInterface[model.PasswordResetToken]
	FindByToken(token string) (*model.PasswordResetToken, error)
}
