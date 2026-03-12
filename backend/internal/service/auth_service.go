package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/email"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/repository/interface"
	_service_interface "github.com/iacopoGhilardi/amILate/internal/service/interface"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"gorm.io/gorm"
)

const (
	resetTokenExpiry     = 15 * time.Minute
	resetPasswordBaseURL = "https://tuaapp.com/reset-password" // sostituisci con il tuo frontend URL
)

type AuthService struct {
	repo           _interface.UserRepositoryInterface
	tokenRepo      _interface.PasswordResetTokenRepositoryInterface
	emailService   _service_interface.EmailServiceInterface
	templateEngine *email.TemplateEngine
}

func NewAuthService(
	userRepo _interface.UserRepositoryInterface,
	tokenRepo _interface.PasswordResetTokenRepositoryInterface,
	emailService _service_interface.EmailServiceInterface,
	templateEngine *email.TemplateEngine,
) *AuthService {
	return &AuthService{
		repo:           userRepo,
		tokenRepo:      tokenRepo,
		emailService:   emailService,
		templateEngine: templateEngine,
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
		ExpiresAt: time.Now().Add(24 * time.Hour),
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
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, nil
}

func (s *AuthService) ForgotPassword(forgotDto dto.ForgotPasswordDto) error {
	logger.Info("Forgot password request for: " + forgotDto.Email)

	user, err := s.repo.FindByEmail(forgotDto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	token, err := generateSecureToken()
	if err != nil {
		return err
	}

	resetToken := &model.PasswordResetToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(resetTokenExpiry),
	}
	if err := s.tokenRepo.Save(resetToken); err != nil {
		return err
	}

	// Render email template
	resetURL := resetPasswordBaseURL + "?token=" + token
	html, err := s.templateEngine.Render(email.TemplateResetPasswordEmail, map[string]string{
		"ResetURL": resetURL,
	})
	if err != nil {
		return err
	}

	// Send email
	return s.emailService.SendResetPasswordEmail(user.Email, html)
}

func (s *AuthService) ResetPassword(resetDto dto.ResetPasswordDto) error {
	logger.Info("Reset password request")

	if resetDto.NewPassword != resetDto.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	resetToken, err := s.tokenRepo.FindByToken(resetDto.Token)
	if err != nil {
		return errors.New("token non valido")
	}
	if resetToken.IsExpired() {
		return errors.New("token scaduto")
	}
	if resetToken.IsUsed() {
		return errors.New("token già utilizzato")
	}

	user, err := s.repo.Find(resetToken.UserID)
	if err != nil {
		return err
	}
	hashedPassword, err := security.HashPassword(resetDto.NewPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := s.repo.Update(user); err != nil {
		return err
	}
	
	now := time.Now()
	resetToken.UsedAt = &now
	return s.tokenRepo.Update(resetToken)
}

func generateSecureToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
