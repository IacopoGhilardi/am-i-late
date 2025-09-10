package test

import (
	"github.com/iacopoGhilardi/amILate/internal/user"
	"github.com/iacopoGhilardi/amILate/pkg/security"
)

type MockUserRepo struct {
	users map[uint]*user.User
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{
		users: make(map[uint]*user.User),
	}
}

type FakeUserRepo struct{}

func (f *FakeUserRepo) Save(u *user.User) error { return nil }
func (f *FakeUserRepo) Find(id uint) (*user.User, error) {
	mockPassword, _ := security.HashPassword("test_password")
	return &user.User{Email: "fake", Password: mockPassword}, nil
}
