package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"github.com/iacopoGhilardi/amILate/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type AuthServiceTestSuite struct {
	suite.Suite
	mockRepo    *mocks.UserRepositoryInterface
	authService *service.AuthService
}

func (s *AuthServiceTestSuite) SetupTest() {
	s.mockRepo = new(mocks.UserRepositoryInterface)
	s.authService = service.NewAuthService(s.mockRepo)
}

func (s *AuthServiceTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func (s *AuthServiceTestSuite) Test_should_register_user_successfully() {
	registerDto := dto.RegistrationDto{
		Email:           "newuser@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Name:            "New User",
	}

	s.mockRepo.On("EmailExists", registerDto.Email).Return(false, nil)

	s.mockRepo.On("Save", mock.MatchedBy(func(u *model.User) bool {
		return u.Email == registerDto.Email &&
			u.Name == registerDto.Name &&
			u.Password != "password123" && // Password deve essere hashata
			len(u.Password) > 0 &&
			u.PublicID != uuid.Nil // PublicID deve essere generato
	})).Return(nil)

	response, err := s.authService.Register(registerDto)

	s.NoError(err)
	s.NotNil(response)
	s.NotEmpty(response.Token)
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_email_already_exists() {
	registerDto := dto.RegistrationDto{
		Email:           "existing@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Name:            "User",
	}

	s.mockRepo.On("EmailExists", registerDto.Email).Return(true, nil)

	response, err := s.authService.Register(registerDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "already registered")
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_passwords_do_not_match() {
	registerDto := dto.RegistrationDto{
		Email:           "user@example.com",
		Password:        "password123",
		ConfirmPassword: "differentPassword",
		Name:            "User",
	}

	s.mockRepo.On("EmailExists", registerDto.Email).Return(false, nil)
	response, err := s.authService.Register(registerDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "do not match")
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_email_check_fails() {
	registerDto := dto.RegistrationDto{
		Email:           "user@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Name:            "User",
	}

	s.mockRepo.On("EmailExists", registerDto.Email).Return(false, errors.New("database error"))
	response, err := s.authService.Register(registerDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "database")
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_save_fails() {
	registerDto := dto.RegistrationDto{
		Email:           "user@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Name:            "User",
	}

	s.mockRepo.On("EmailExists", registerDto.Email).Return(false, nil)

	s.mockRepo.On("Save", mock.Anything).Return(errors.New("save failed"))
	response, err := s.authService.Register(registerDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "save failed")
}

func (s *AuthServiceTestSuite) Test_should_login_user_successfully() {
	loginDto := dto.LoginDto{
		Email:    "user@example.com",
		Password: "password123",
	}

	hashedPassword, _ := security.HashPassword("password123")

	existingUser := &model.User{
		Email:    loginDto.Email,
		Password: hashedPassword,
		PublicID: uuid.New(),
		Name:     "Test User",
	}

	s.mockRepo.On("FindByEmail", loginDto.Email).Return(existingUser, nil)
	response, err := s.authService.Login(loginDto)

	s.NoError(err)
	s.NotNil(response)
	s.NotEmpty(response.Token)
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_user_not_found_on_login() {
	loginDto := dto.LoginDto{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	s.mockRepo.On("FindByEmail", loginDto.Email).Return(nil, gorm.ErrRecordNotFound)
	response, err := s.authService.Login(loginDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "invalid email or password")
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_password_is_incorrect() {
	loginDto := dto.LoginDto{
		Email:    "user@example.com",
		Password: "wrongPassword",
	}

	hashedPassword, _ := security.HashPassword("correctPassword")
	existingUser := &model.User{
		Email:    loginDto.Email,
		Password: hashedPassword,
		PublicID: uuid.New(),
		Name:     "Test User",
	}

	s.mockRepo.On("FindByEmail", loginDto.Email).Return(existingUser, nil)
	response, err := s.authService.Login(loginDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "invalid email or password")
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_database_fails_on_login() {
	loginDto := dto.LoginDto{
		Email:    "user@example.com",
		Password: "password123",
	}

	s.mockRepo.On("FindByEmail", loginDto.Email).Return(nil, errors.New("database error"))
	response, err := s.authService.Login(loginDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "database")
}

func (s *AuthServiceTestSuite) Test_should_return_error_when_user_has_empty_password() {
	loginDto := dto.LoginDto{
		Email:    "user@example.com",
		Password: "password123",
	}

	existingUser := &model.User{
		Email:    loginDto.Email,
		Password: "",
		PublicID: uuid.New(),
		Name:     "Test User",
	}

	s.mockRepo.On("FindByEmail", loginDto.Email).Return(existingUser, nil)
	response, err := s.authService.Login(loginDto)

	s.Error(err)
	s.Nil(response)
	s.Contains(err.Error(), "invalid email or password")
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
