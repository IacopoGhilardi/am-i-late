package alarm

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

type Alarm struct {
	commons.BaseModelWithSafeDelete

	Time string `gorm:"column:time;type:varchar(10);not null"`
	//UserID        uint      `gorm:"column:user_id;not null;index"`
	DestinationID *uint     `gorm:"column:destination_id;index"`
	Active        bool      `gorm:"column:active;default:true"`
	Label         string    `gorm:"column:label;type:varchar(255)"`
	PublicID      uuid.UUID `gorm:"column:public_id;type:varchar(255);unique"`
}
