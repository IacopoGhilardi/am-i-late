package service

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/repository"
)

type AppointmentService struct {
	repo repository.AppointmentRepositoryInterface
}

func NewAppointmentService(repo repository.AppointmentRepositoryInterface) *AppointmentService {
	return &AppointmentService{
		repo: repo,
	}
}

func (s *AppointmentService) GetAllAppointments() ([]model.Appointment, error) {
	return s.repo.FindAll()
}

func (s *AppointmentService) GetAppointmentByID(id uint) (*model.Appointment, error) {
	return s.repo.Find(id)
}

func (s *AppointmentService) GetAppointmentByPublicId(uuid uuid.UUID) (*model.Appointment, error) {
	return s.repo.FindByPublicId(uuid)
}

func (s *AppointmentService) CreateAppointment(a *model.Appointment) (*model.Appointment, error) {
	err := s.repo.Save(a)
	return a, err
}

func (s *AppointmentService) DeleteAppointment(id uint) error {
	return s.repo.Delete(id)
}

func (s *AppointmentService) DeleteAppointmentFromPublicId(publicId uuid.UUID) error {
	app, err := s.repo.FindByPublicId(publicId)
	if err != nil {
		return err
	}
	return s.repo.Delete(app.ID)
}
