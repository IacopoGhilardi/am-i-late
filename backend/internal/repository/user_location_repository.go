package repository

import (
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

type UserLocationRepository struct {
	*BaseRepository[model.UserLocation]
}

func NewUserLocationRepository() *UserLocationRepository {
	return &UserLocationRepository{
		BaseRepository: NewBaseRepository[model.UserLocation](db.GetDB()),
	}
}

func (r *UserLocationRepository) Upsert(location *model.UserLocation) error {
	var existing model.UserLocation
	result := r.Db.Where("user_id = ?", location.UserID).First(&existing)

	if result.Error != nil {
		return r.Db.Create(location).Error
	}

	existing.Latitude = location.Latitude
	existing.Longitude = location.Longitude
	return r.Db.Save(&existing).Error
}
