package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

type Destination struct {
	commons.BaseModel

	PublicID          uuid.UUID         `gorm:"column:public_id;type:uuid;unique;not null"`
	UserID            uint              `gorm:"column:user_id;not null;index"`
	FullAddress       string            `gorm:"column:full_address;type:varchar(255);not null"`
	Name              string            `gorm:"column:name;type:varchar(255);not null"`
	AddressComponents AddressComponents `gorm:"embedded"`
	Location          commons.Location  `gorm:"embedded"`
	TimeZone          string            `gorm:"column:time_zone;type:varchar(255);not null"`
	appointments      []Appointment     `gorm:"foreignKey:DestinationId"`
	Metadata          Metadata
}

func (Destination) TableName() string {
	return "destinations"
}

type AddressComponents struct {
	FormattedAddress string `gorm:"column:formatted_address;type:varchar(255);not null"`
	GooglePlaceId    string `gorm:"column:google_place_id;type:varchar(255);not null"`
}

type Metadata struct {
	IsSaved     bool          `gorm:"column:is_saved;default:false"`
	LastUsedAt  *time.Time    `gorm:"column:last_used_at;index"`
	DeleteAfter time.Duration `gorm:"column:delete_after;type:varchar(255);not null"`
}
