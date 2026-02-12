package service

import (
	"errors"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	"github.com/iacopoGhilardi/amILate/internal/repository/interface"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"gorm.io/gorm"
)

type AuthService struct {
	repo _interface.UserRepositoryInterface
}

func NewAuthService(userRepo _interface.UserRepositoryInterface) *AuthService {
	return &AuthService{
		repo: userRepo,
	}
}

func (s *AuthService) Register(registerDto dto.RegistrationDto) (*dto.LoginResponseDto, error) {
	logger.Info("Registering user")
	exists, err := s.repo.EmailExists(registerDto.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		logger.Warn("Email already registered")
		return nil, errors.New("email already registered")
	}

	if registerDto.Password != registerDto.ConfirmPassword {
		logger.Warn("Passwords do not match")
		return nil, errors.New(
			"passwords do not match",
		)
	}
	user := mapper.MapFromRegistrationDto(registerDto)
	hashedPassword, err := security.HashPassword(registerDto.Password)
	if err != nil {
		logger.Error("Error hashing password: " + err.Error())
		return nil, err
	}
	user.Password = hashedPassword

	err = s.repo.Save(user)
	if err != nil {
		logger.Error("Error saving user: " + err.Error())
		return nil, err
	}

	token, err := security.GenerateJWT(user.PublicID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		Token:     token,
		ExpiresAt: 0,
	}, nil
}

func (s *AuthService) Login(loginDto dto.LoginDto) (*dto.LoginResponseDto, error) {
	user, err := s.repo.FindByEmail(loginDto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !security.CheckPasswordHash(loginDto.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := security.GenerateJWT(user.PublicID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		Token: token,
	}, nil
}
