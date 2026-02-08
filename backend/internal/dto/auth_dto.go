package dto

type RegistrationDto struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	ConfirmPassword    string `json:"confirm_password"`
	Name               string `json:"name"`
	AgeConfirmed       bool
	PrivacyAccepted    bool
	TermsAccepted      bool
	LocationPermission bool
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Token     string `json:"token"`
	ExpiresAt int    `json:"expires_at"`
}
