package dto

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
	Token     string `json:"token"`
	ExpiresAt int    `json:"expires_at"`
}
