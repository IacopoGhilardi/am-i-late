package test

import (
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/pkg/security"
)

type MockUserRepo struct {
	users map[uint]*model.User
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{
		users: make(map[uint]*model.User),
	}
}

type FakeUserRepo struct{}

func (f *FakeUserRepo) Save(u *model.User) error { return nil }
func (f *FakeUserRepo) Find(id uint) (*model.User, error) {
	mockPassword, _ := security.HashPassword("test_password")
	return &model.User{Email: "fake", Password: mockPassword}, nil
}
