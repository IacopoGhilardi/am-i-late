package service

import (
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/repository"
)

type DestinationService struct {
	repo *repository.DestinationRepository
}

func NewDestinationService() *DestinationService {
	return &DestinationService{
		repo: repository.NewDestinationRepository(),
	}
}

func (s *DestinationService) GetAllDestinations() ([]model.Destination, error) {
	return s.repo.FindAll()
}

func (s *DestinationService) GetDestinationByID(id uint) (*model.Destination, error) {
	return s.repo.Find(id)
}

func (s *DestinationService) CreateDestination(d *model.Destination) (*model.Destination, error) {
	err := s.repo.Save(d)
	return d, err
}

func (s *DestinationService) DeleteDestination(id uint) error {
	return s.repo.Delete(id)
}
