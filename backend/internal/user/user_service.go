package user

import "github.com/iacopoGhilardi/amILate/pkg/security"

type UserService struct {
	repo *UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: NewUserRepository(),
	}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.repo.Find(id)
}

func (s *UserService) CreateUser(u *User) (*User, error) {
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
