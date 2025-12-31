package route

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/destination"
)

type Route struct {
	commons.BaseModelWithSafeDelete

	PublicID    uuid.UUID               `gorm:"column:public_id;type:uuid;unique;not null"`
	UserID      uint                    `gorm:"column:user_id;not null;index"`
	Origin      Origin                  `Gorm:"embedded"`
	Destination destination.Destination `gorm:"foreignKey:DestinationID"`
	MetaData    MetaData                `gorm:"embedded"`
	IsSingleUse bool                    `gorm:"column:is_single_use"`
}

type Origin struct {
	Type     string           `gorm:"column:type;type:varchar(255);not null"`
	Location commons.Location `gorm:"embedded"`
}

type MetaData struct {
	RouteProviderId     string `gorm:"column:route_provider_id;type:varchar(255);not null"`
	AverageDuration     int    `gorm:"column:average_duration"`
	LastDurationSeconds int    `gorm:"column:last_duration_seconds"`
	TrafficIntensity    int    `gorm:"column:traffic_intensity"`
}
