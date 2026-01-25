package repository

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type DestinationRepository struct {
	*commons.BaseRepository[model.Destination]
}

func NewDestinationRepository() *DestinationRepository {
	return &DestinationRepository{
		BaseRepository: commons.NewBaseRepository[model.Destination](db.GetDB()),
	}
}
