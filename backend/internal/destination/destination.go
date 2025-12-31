package destination

import (
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/alarm"
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

var availableTransportModes = []string{"car", "bus", "train", "plane"}

type Destination struct {
	commons.BaseModelWithSafeDelete

	PublicID          uuid.UUID         `gorm:"column:public_id;type:uuid;unique;not null"`
	UserID            uint              `gorm:"column:user_id;not null;index"`
	FullAddress       string            `gorm:"column:full_address;type:varchar(255);not null"`
	Name              string            `gorm:"column:name;type:varchar(255);not null"`
	AddressComponents AddressComponents `gorm:"embedded"`
	Location          commons.Location  `gorm:"embedded"`
	TransportMode     string            `gorm:"column:transport_mode;type:varchar(255);not null"`
	TimeZone          string            `gorm:"column:time_zone;type:varchar(255);not null"`
	Alarms            []alarm.Alarm     `gorm:"foreignKey:DestinationID"`
}

type AddressComponents struct {
	FormattedAddress string `gorm:"column:formatted_address;type:varchar(255);not null"`
	GooglePlaceId    string `gorm:"column:google_place_id;type:varchar(255);not null"`
}

type Metadata struct {
	VisitCount    int        `gorm:"column:visit_count"`
	LastVisitedAt *time.Time `gorm:"column:last_visited_at;index"`
	IsFavorite    bool       `gorm:"column:is_favorite"`
}

func (d *Destination) ValidateTransport(transportMode string) bool {
	return slices.Contains(availableTransportModes, transportMode)
}
