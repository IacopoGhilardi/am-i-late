package dto

import "time"

type RegistrationDto struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	ConfirmPassword    string `json:"confirm_password"`
	Name               string `json:"name"`
	AgeConfirmed       bool   `json:"age_confirmed"`
	PrivacyAccepted    bool   `json:"privacy_accepted"`
	TermsAccepted      bool   `json:"terms_accepted"`
	LocationPermission bool   `json:"location_permission"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type ForgotPasswordDto struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordDto struct {
	Token           string `json:"token" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}
