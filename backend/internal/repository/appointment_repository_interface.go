package repository

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type AppointmentRepositoryInterface interface {
	BaseRepositoryInterface[model.Appointment] // Eredita tutti i metodi dell'interfaccia BaseRepository

	FindByPublicId(publicId uuid.UUID) (*model.Appointment, error)
}
