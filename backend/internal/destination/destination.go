package destination

import (
	"github.com/iacopoGhilardi/amILate/internal/alarm"
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

type Destination struct {
	commons.BaseModelWithSafeDelete

	Name       string        `gorm:"column:name;type:varchar(255);not null"`
	Address    string        `gorm:"column:address;type:varchar(255);not null"`
	City       string        `gorm:"column:city;type:varchar(255);not null"`
	PostalCode string        `gorm:"column:postal_code;type:varchar(20)"`
	Country    string        `gorm:"column:country;type:varchar(100);not null"`
	Latitude   float64       `gorm:"column:latitude;type:decimal(9,6)"`
	Longitude  float64       `gorm:"column:longitude;type:decimal(9,6)"`
	UserID     uint          `gorm:"column:user_id;not null;index"`
	Alarms     []alarm.Alarm `gorm:"foreignKey:DestinationID"`
}
