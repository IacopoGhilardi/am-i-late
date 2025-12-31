package alarm

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/db"
)

type AlarmRepository struct {
	*commons.BaseRepository[Alarm]
}

func NewAlarmRepository() *AlarmRepository {
	return &AlarmRepository{
		BaseRepository: commons.NewBaseRepository[Alarm](db.GetDB()),
	}
}

func (r *AlarmRepository) FindActiveByUser(userID uint) ([]Alarm, error) {
	var alarms []Alarm
	result := db.GetDB().Where("user_id = ? AND active = ?", userID, true).Find(&alarms)
	return alarms, result.Error
}

func (r *AlarmRepository) FindByPublicId(id uuid.UUID) (*Alarm, error) {
	var alarm Alarm
	result := db.GetDB().Where("public_id = ?", id).First(&alarm)
	return &alarm, result.Error
}
