package dto

import (
	"github.com/google/uuid"
)

type AppointmentDto struct {
	ID                     uuid.UUID      `json:"id"`
	Destination            DestinationDto `json:"destination"`
	ScheduledAt            string         `json:"scheduled_at"`
	TransportMode          string         `json:"transport_mode"`
	Status                 string         `json:"status"`
	EstimatedTravelMinutes int            `json:"estimated_travel_minutes"`
}

type CreateAppointmentRequestDto struct {
	Destination   CreateDestinationRequestDto `json:"destination"`
	ScheduledAt   string                      `json:"scheduled_at"`
	TransportMode string                      `json:"transport_mode"`
}
