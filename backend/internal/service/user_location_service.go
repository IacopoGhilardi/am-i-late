package service

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
	_interface "github.com/iacopoGhilardi/amILate/internal/repository/interface"
)

type UserLocationService struct {
	repo     _interface.UserLocationRepositoryInterface
	userRepo _interface.UserRepositoryInterface
}

func NewUserLocationService(repo _interface.UserLocationRepositoryInterface,
	userRepo _interface.UserRepositoryInterface) *UserLocationService {
	return &UserLocationService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *UserLocationService) UpdateLocation(publicID uuid.UUID, lat, lng float64) error {
	user, err := s.userRepo.FindByPublicId(publicID)
	if err != nil {
		return err
	}

	return s.repo.Upsert(&model.UserLocation{
		UserID:    user.ID,
		Latitude:  lat,
		Longitude: lng,
	})
}
