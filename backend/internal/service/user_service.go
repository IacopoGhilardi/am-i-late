package service

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.Find(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetUserByPublicId(publicId uuid.UUID) (*model.User, error) {
	return s.repo.FindByPublicId(publicId)
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
