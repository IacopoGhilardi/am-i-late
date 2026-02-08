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
