package mapper

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

func MapUserToDto(user model.User) *dto.UserDto {
	return &dto.UserDto{
		Email:              user.Email,
		Name:               user.Name,
		AgeConfirmed:       user.AgeConfirmed,
		PrivacyAccepted:    user.PrivacyAccepted,
		TermsAccepted:      user.TermsAccepted,
		LocationPermission: user.LocationPermission,
	}
}

func MapUserFromDto(dto dto.UserDto) *model.User {
	return &model.User{
		Email:              dto.Email,
		Name:               dto.Name,
		AgeConfirmed:       dto.AgeConfirmed,
		PrivacyAccepted:    dto.PrivacyAccepted,
		TermsAccepted:      dto.TermsAccepted,
		LocationPermission: dto.LocationPermission,
		PublicID:           dto.Id,
	}
}

func MapFromRegistrationDto(dto dto.RegistrationDto) *model.User {
	return &model.User{
		Email:              dto.Email,
		Name:               dto.Name,
		AgeConfirmed:       dto.AgeConfirmed,
		PrivacyAccepted:    dto.PrivacyAccepted,
		TermsAccepted:      dto.TermsAccepted,
		LocationPermission: dto.LocationPermission,
		PublicID:           uuid.New(),
	}
}
