package repository

import (
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type DestinationRepository struct {
	*BaseRepository[model.Destination]
}

func NewDestinationRepository() *DestinationRepository {
	return &DestinationRepository{
		BaseRepository: NewBaseRepository[model.Destination](db.GetDB()),
	}
}
