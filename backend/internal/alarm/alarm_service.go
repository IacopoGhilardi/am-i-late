package alarm

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

func (s *AlarmService) GetAlarmByID(id uint) (*Alarm, error) {
	return s.repo.Find(id)
}

func (s *AlarmService) CreateAlarm(a *Alarm) (*Alarm, error) {
	err := s.repo.Save(a)
	return a, err
}

func (s *AlarmService) DeleteAlarm(id uint) error {
	return s.repo.Delete(id)
}

func (s *AlarmService) GetActiveAlarmsByUser(userID uint) ([]Alarm, error) {
	return s.repo.FindActiveByUser(userID)
}
