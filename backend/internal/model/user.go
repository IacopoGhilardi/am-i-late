package model

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

type User struct {
	commons.BaseModel
	Email              string    `gorm:"column:email;type:varchar(255);unique;not null"`
	Password           string    `gorm:"column:password;type:varchar(255);not null"`
	Name               string    `gorm:"column:name;type:varchar(255);not null"`
	AgeConfirmed       bool      `gorm:"column:age_confirmed;type:boolean;default:false"`
	PrivacyAccepted    bool      `gorm:"column:privacy_accepted;type:boolean;default:false"`
	TermsAccepted      bool      `gorm:"column:terms_accepted;type:boolean;default:false"`
	LocationPermission bool      `gorm:"column:location_permission;type:boolean;default:false"`
	PublicID           uuid.UUID `gorm:"column:public_id;type:uuid;unique;not null"`
}
