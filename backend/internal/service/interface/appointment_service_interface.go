package _interface

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type AppointmentServiceInterface interface {
	GetAllAppointments() ([]model.Appointment, error)
	GetAppointmentByID(id uint) (*model.Appointment, error)
	GetAppointmentByPublicId(uuid uuid.UUID) (*model.Appointment, error)
	CreateAppointment(a *model.Appointment) (*model.Appointment, error)
	DeleteAppointment(id uint) error
	DeleteAppointmentFromPublicId(publicId uuid.UUID) error
}
