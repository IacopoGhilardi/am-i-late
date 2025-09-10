package destination

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/db"
)

type DestinationRepository struct {
	*commons.BaseRepository[Destination]
}

func NewDestinationRepository() *DestinationRepository {
	return &DestinationRepository{
		BaseRepository: commons.NewBaseRepository[Destination](db.GetDB()),
	}
}
