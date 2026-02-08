package repository

import "github.com/iacopoGhilardi/amILate/internal/model"

type DestinationRepositoryInterface interface {
	BaseRepositoryInterface[model.Destination]
}
