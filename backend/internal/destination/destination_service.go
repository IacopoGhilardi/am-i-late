package destination

type DestinationService struct {
	repo *DestinationRepository
}

func NewDestinationService() *DestinationService {
	return &DestinationService{
		repo: NewDestinationRepository(),
	}
}

func (s *DestinationService) GetAllDestinations() ([]Destination, error) {
	return s.repo.FindAll()
}

func (s *DestinationService) GetDestinationByID(id uint) (*Destination, error) {
	return s.repo.Find(id)
}

func (s *DestinationService) CreateDestination(d *Destination) (*Destination, error) {
	err := s.repo.Save(d)
	return d, err
}

func (s *DestinationService) DeleteDestination(id uint) error {
	return s.repo.Delete(id)
}
