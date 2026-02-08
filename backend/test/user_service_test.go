package test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"github.com/iacopoGhilardi/amILate/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var userEmailTest = "test@example.com"

// Suite contiene lo stato per ogni test
type UserServiceTestSuite struct {
	suite.Suite
	mockRepo    *mocks.UserRepositoryInterface
	userService *service.UserService
}

func (s *UserServiceTestSuite) SetupTest() {
	s.mockRepo = new(mocks.UserRepositoryInterface)
	s.userService = service.NewUserService(s.mockRepo)
}

func (s *UserServiceTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func (s *UserServiceTestSuite) Test_should_find_user_by_public_id() {
	testUUID := uuid.New()
	testUser := &model.User{
		Email:              userEmailTest,
		Name:               "test_name",
		AgeConfirmed:       false,
		PrivacyAccepted:    false,
		TermsAccepted:      false,
		LocationPermission: false,
		PublicID:           testUUID,
	}

	s.mockRepo.On("FindByPublicId", testUUID).Return(testUser, nil)

	res, err := s.userService.GetUserByPublicId(testUUID)
	s.NoError(err)
	s.Equal(testUser.PublicID, res.PublicID)
	s.Equal(testUser.Email, res.Email)
}

func (s *UserServiceTestSuite) Test_should_not_find_user_by_public_id() {
	testUUID := uuid.New()
	s.mockRepo.On("FindByPublicId", testUUID).Return(nil, nil)

	res, err := s.userService.GetUserByPublicId(testUUID)

	s.NoError(err)
	s.Nil(res)
}

func (s *UserServiceTestSuite) Test_Should_Return_Error_When_UUID_Is_Invalid() {
	invalidUUID := uuid.Nil

	s.mockRepo.On("FindByPublicId", invalidUUID).
		Return(nil, errors.New("invalid uuid"))

	res, err := s.userService.GetUserByPublicId(invalidUUID)
	s.Error(err)
	s.Nil(res)
	s.Contains(err.Error(), "invalid")
}

func (s *UserServiceTestSuite) Test_should_find_user_by_email() {
	testUUID := uuid.New()
	testUser := &model.User{
		Email:              userEmailTest,
		Name:               "test_name",
		AgeConfirmed:       false,
		PrivacyAccepted:    false,
		TermsAccepted:      false,
		LocationPermission: false,
		PublicID:           testUUID,
	}

	s.mockRepo.On("FindByEmail", userEmailTest).Return(testUser, nil)
	res, err := s.userService.GetUserByEmail(userEmailTest)

	s.NoError(err)
	s.Equal(testUser.Email, res.Email)
	s.Equal(testUser.PublicID, res.PublicID)
}

func (s *UserServiceTestSuite) Test_should_not_find_user_by_email() {
	s.mockRepo.On("FindByEmail", userEmailTest).Return(nil, nil)
	res, err := s.userService.GetUserByEmail(userEmailTest)

	s.NoError(err)
	s.Nil(res)
}

func (s *UserServiceTestSuite) Test_should_create_user_successfully() {
	newUser := &model.User{
		Email:              "newuser@example.com",
		Password:           "plainPassword123",
		Name:               "New User",
		AgeConfirmed:       true,
		PrivacyAccepted:    true,
		TermsAccepted:      true,
		LocationPermission: true,
	}

	s.mockRepo.On("Save", mock.MatchedBy(func(u *model.User) bool {
		return u.Email == newUser.Email &&
			u.Password != "plainPassword123" &&
			len(u.Password) > 0
	})).Return(nil)

	createdUser, err := s.userService.CreateUser(newUser)

	s.NoError(err)
	s.NotNil(createdUser)
	s.Equal("newuser@example.com", createdUser.Email)
	s.NotEqual("plainPassword123", createdUser.Password)
	hashedPassword, _ := security.HashPassword("plainPassword123")
	s.NotEqual(hashedPassword, createdUser.Password)
	s.NotEmpty(createdUser.Password)
}

func (s *UserServiceTestSuite) Test_should_return_error_when_save_fails() {
	newUser := &model.User{
		Email:    "newuser@example.com",
		Password: "plainPassword123",
		Name:     "New User",
	}

	s.mockRepo.On("Save", mock.Anything).Return(errors.New("database error"))
	createdUser, err := s.userService.CreateUser(newUser)

	s.Error(err)
	s.NotNil(createdUser)
	s.Contains(err.Error(), "database")
}

func (s *UserServiceTestSuite) Test_should_delete_user_successfully() {
	userID := uint(123)
	s.mockRepo.On("Delete", userID).Return(nil)

	err := s.userService.DeleteUser(userID)
	s.NoError(err)
}

func (s *UserServiceTestSuite) Test_should_return_error_when_delete_fails() {
	userID := uint(456)

	s.mockRepo.On("Delete", userID).Return(errors.New("database error"))
	err := s.userService.DeleteUser(userID)

	s.Error(err)
	s.Contains(err.Error(), "database")
}

func (s *UserServiceTestSuite) Test_should_return_error_when_user_not_found_on_delete() {
	userID := uint(999)

	s.mockRepo.On("Delete", userID).Return(errors.New("record not found"))

	err := s.userService.DeleteUser(userID)

	s.Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *UserServiceTestSuite) Test_should_not_delete_user_with_zero_id() {
	userID := uint(0)

	s.mockRepo.On("Delete", userID).Return(errors.New("invalid id"))

	err := s.userService.DeleteUser(userID)

	s.Error(err)
	s.Contains(err.Error(), "invalid")
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
