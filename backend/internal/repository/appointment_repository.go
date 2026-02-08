package repository

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type AppointmentRepository struct {
	*BaseRepository[model.Appointment]
}

func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{
		BaseRepository: NewBaseRepository[model.Appointment](db.GetDB()),
	}
}

func (r *AppointmentRepository) FindByPublicId(publicId uuid.UUID) (*model.Appointment, error) {
	var apt model.Appointment
	err := r.Db.Where("public_id = ?", publicId).First(&apt).Error
	if err != nil {
		return nil, err
	}
	return &apt, nil
}
