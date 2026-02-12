package _interface

import "github.com/iacopoGhilardi/amILate/internal/model"

type DestinationServiceInterface interface {
	GetAllDestinations() ([]model.Destination, error)
	GetDestinationByID(id uint) (*model.Destination, error)
	CreateDestination(d *model.Destination) (*model.Destination, error)
	DeleteDestination(id uint) error
}
