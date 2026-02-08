package service

import (
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/repository"
)

type DestinationService struct {
	repo repository.DestinationRepositoryInterface
}

func NewDestinationService(repo repository.DestinationRepositoryInterface) *DestinationService {
	return &DestinationService{
		repo: repo,
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

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *DestinationService) DeleteDestination(id uint) error {
	return s.repo.Delete(id)
}
