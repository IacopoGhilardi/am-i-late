package mapper

import (
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

func MapAppointmentToDto(appointment model.Appointment) dto.AppointmentDto {
	return dto.AppointmentDto{
		ID:                     appointment.PublicId,
		Destination:            *MapDestinationToDto(appointment.Destination),
		EstimatedTravelMinutes: appointment.EstimatedTravelMinutes,
		ScheduledAt:            appointment.ScheduledAt,
		Status:                 appointment.Status,
		TransportMode:          appointment.TransportMode,
	}
}

func MapFromCreateAppointmentRequest(request dto.CreateAppointmentRequestDto) model.Appointment {
	return model.Appointment{
		TransportMode: request.TransportMode,
		ScheduledAt:   request.ScheduledAt,
	}
}
