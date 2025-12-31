package alarm

import "github.com/google/uuid"

type AlarmService struct {
	repo *AlarmRepository
}

func NewAlarmService() *AlarmService {
	return &AlarmService{
		repo: NewAlarmRepository(),
	}
}

func (s *AlarmService) GetAllAlarms() ([]Alarm, error) {
	return s.repo.FindAll()
}

func (s *AlarmService) GetAlarmByID(id uuid.UUID) (*Alarm, error) {
	return s.repo.FindByPublicId(id)
}

func (s *AlarmService) CreateAlarm(a *CreateAlarmRequestDto) (*Alarm, error) {
	alarm := MapFromCreateReq(*a)
	err := s.repo.Save(alarm)
	return alarm, err
}

func (s *AlarmService) DeleteAlarm(id uint) error {
	return s.repo.Delete(id)
}

func (s *AlarmService) GetActiveAlarmsByUser(userID uint) ([]Alarm, error) {
	return s.repo.FindActiveByUser(userID)
}
