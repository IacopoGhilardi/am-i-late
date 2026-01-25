package dto

import "github.com/google/uuid"

type UserDto struct {
	Email              string
	Name               string
	Id                 uuid.UUID
	AgeConfirmed       bool
	PrivacyAccepted    bool
	TermsAccepted      bool
	LocationPermission bool
}

type RegistrationDto struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
