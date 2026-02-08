package model

import (
	"slices"

	"github.com/google/uuid"
)

var transportModes = []string{"car", "public_transport", "foot"}
var DefaultTransportMode = transportModes[0]
var notificationStates = []string{"pending", "monitoring", "sent", "cancelled"}
var statuses = []string{"scheduled", "completed", "cancelled"}

type Appointment struct {
	BaseModel

	DestinationID          uint        `gorm:"column:destination_id;not null;index"`
	Destination            Destination `gorm:"foreignKey:DestinationID"`
	UserID                 uint        `gorm:"column:user_id;not null;index"`
	PublicId               uuid.UUID   `gorm:"column:public_id;type:varchar(255);unique;not null"`
	ScheduledAt            string      `gorm:"column:scheduled_at;type:varchar(255);not null"`
	TransportMode          string      `gorm:"column:transport_mode;type:varchar(255);not null"`
	EstimatedTravelMinutes int         `gorm:"column:estimated_travel_minutes;type:bigint;not null"`
	EstimatedTravelRange   int         `gorm:"column:estimated_travel_range;type:bigint;not null"`
	LastTravelUpdateAt     string      `gorm:"column:last_travel_update_at;type:varchar(255);not null"`
	NotificationState      string      `gorm:"column:notification_state;type:varchar(255);not null"`
	Status                 string      `gorm:"column:status;type:varchar(255);not null"`
	DeleteAfter            string      `gorm:"column:delete_after;type:varchar(255);not null"`
	GeoFenceId             string      `gorm:"column:geo_fence_id;type:varchar(255);not null"`
}

func (Appointment) TableName() string {
	return "appointments"
}

func (a Appointment) ValidateTransport(transport string) bool {
	return slices.Contains(transportModes, transport)
}

func (a Appointment) ValidateStatus(status string) bool {
	return slices.Contains(statuses, status)
}

func (a Appointment) ValidateNotificationState(state string) bool {
	return slices.Contains(notificationStates, state)
}
