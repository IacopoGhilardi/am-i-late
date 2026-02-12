package _interface

import "github.com/iacopoGhilardi/amILate/internal/dto"

type AuthServiceInterface interface {
	Register(registerDto dto.RegistrationDto) (*dto.LoginResponseDto, error)
	Login(loginDto dto.LoginDto) (*dto.LoginResponseDto, error)
}
