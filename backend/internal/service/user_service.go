package service

import (
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/pkg/security"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.Find(id)
}

func (s *UserService) CreateUser(u *model.User) (*model.User, error) {
	hashedPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword

	err = s.repo.Save(u)
	return u, err
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
