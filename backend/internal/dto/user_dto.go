package dto

import "github.com/google/uuid"

type UserDto struct {
	Email              string    `json:"email" validate:"required,email"`
	Name               string    `json:"name" validate:"required,max=255"`
	Id                 uuid.UUID `json:"id" validate:"required,uuid"`
	AgeConfirmed       bool      `json:"age_confirmed"`
	PrivacyAccepted    bool      `json:"privacy_accepted"`
	TermsAccepted      bool      `json:"terms_accepted"`
	LocationPermission bool      `json:"location_permission"`
}

type UpdateLocationDto struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}
