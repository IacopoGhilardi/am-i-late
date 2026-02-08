package test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/iacopoGhilardi/amILate/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DestinationServiceTestSuite struct {
	suite.Suite
	mockRepo           *mocks.DestinationRepositoryInterface
	destinationService *service.DestinationService
}

func (s *DestinationServiceTestSuite) SetupTest() {
	s.mockRepo = new(mocks.DestinationRepositoryInterface)
	s.destinationService = service.NewDestinationService(s.mockRepo)
}

func (s *DestinationServiceTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func (s *DestinationServiceTestSuite) Test_should_get_all_destinations_successfully() {
	expectedDestinations := []model.Destination{
		{
			PublicID:    uuid.New(),
			UserID:      1,
			FullAddress: "Via Roma 123, Milano, Italy",
			Name:        "Office",
			AddressComponents: model.AddressComponents{
				FormattedAddress: "Via Roma 123, 20100 Milano MI, Italy",
				GooglePlaceId:    "ChIJAWaem...",
			},
			Location: model.Location{
				Latitude:  45.4642,
				Longitude: 9.1900,
			},
			TimeZone: "Europe/Rome",
			Metadata: model.Metadata{
				IsSaved:     true,
				DeleteAfter: 24 * time.Hour,
			},
		},
		{
			PublicID:    uuid.New(),
			UserID:      1,
			FullAddress: "Via Torino 456, Milano, Italy",
			Name:        "Home",
			AddressComponents: model.AddressComponents{
				FormattedAddress: "Via Torino 456, 20100 Milano MI, Italy",
				GooglePlaceId:    "ChIJBXbfn...",
			},
			Location: model.Location{
				Latitude:  45.4500,
				Longitude: 9.2000,
			},
			TimeZone: "Europe/Rome",
		},
	}

	s.mockRepo.On("FindAll").Return(expectedDestinations, nil)

	destinations, err := s.destinationService.GetAllDestinations()

	s.NoError(err)
	s.NotNil(destinations)
	s.Len(destinations, 2)
	s.Equal("Office", destinations[0].Name)
	s.Equal("Home", destinations[1].Name)
	s.Equal("Via Roma 123, Milano, Italy", destinations[0].FullAddress)
}

func (s *DestinationServiceTestSuite) Test_should_return_empty_list_when_no_destinations() {
	s.mockRepo.On("FindAll").Return([]model.Destination{}, nil)

	destinations, err := s.destinationService.GetAllDestinations()

	s.NoError(err)
	s.NotNil(destinations)
	s.Empty(destinations)
}

func (s *DestinationServiceTestSuite) Test_should_return_error_when_findall_fails() {
	s.mockRepo.On("FindAll").Return(nil, errors.New("database error"))

	destinations, err := s.destinationService.GetAllDestinations()

	s.Error(err)
	s.Nil(destinations)
	s.Contains(err.Error(), "database")
}

func (s *DestinationServiceTestSuite) Test_should_find_destination_by_id() {
	destinationID := uint(123)
	publicID := uuid.New()
	expectedDestination := &model.Destination{
		BaseModel: model.BaseModel{
			ID: destinationID,
		},
		PublicID:    publicID,
		UserID:      1,
		FullAddress: "Via Roma 123, Milano, Italy",
		Name:        "Office",
		AddressComponents: model.AddressComponents{
			FormattedAddress: "Via Roma 123, 20100 Milano MI, Italy",
			GooglePlaceId:    "ChIJAWaem...",
		},
		Location: model.Location{
			Latitude:  45.4642,
			Longitude: 9.1900,
		},
		TimeZone: "Europe/Rome",
		Metadata: model.Metadata{
			IsSaved:     true,
			DeleteAfter: 24 * time.Hour,
		},
	}

	s.mockRepo.On("Find", destinationID).Return(expectedDestination, nil)

	destination, err := s.destinationService.GetDestinationByID(destinationID)

	s.NoError(err)
	s.NotNil(destination)
	s.Equal(destinationID, destination.ID)
	s.Equal(publicID, destination.PublicID)
	s.Equal("Office", destination.Name)
	s.Equal("Via Roma 123, Milano, Italy", destination.FullAddress)
	s.Equal(45.4642, destination.Location.Latitude)
	s.Equal(9.1900, destination.Location.Longitude)
}

func (s *DestinationServiceTestSuite) Test_should_return_nil_when_destination_not_found() {
	destinationID := uint(999)

	s.mockRepo.On("Find", destinationID).Return(nil, nil)

	destination, err := s.destinationService.GetDestinationByID(destinationID)

	s.NoError(err)
	s.Nil(destination)
}

func (s *DestinationServiceTestSuite) Test_should_return_error_when_find_fails() {
	destinationID := uint(123)

	s.mockRepo.On("Find", destinationID).Return(nil, errors.New("database error"))

	destination, err := s.destinationService.GetDestinationByID(destinationID)

	s.Error(err)
	s.Nil(destination)
	s.Contains(err.Error(), "database")
}

func (s *DestinationServiceTestSuite) Test_should_create_destination_successfully() {
	newDestination := &model.Destination{
		PublicID:    uuid.New(),
		UserID:      1,
		FullAddress: "Piazza Duomo, Milano, Italy",
		Name:        "Milan Cathedral",
		AddressComponents: model.AddressComponents{
			FormattedAddress: "Piazza del Duomo, 20122 Milano MI, Italy",
			GooglePlaceId:    "ChIJXQHkbCnGhkcRLGDdkFP5Prc",
		},
		Location: model.Location{
			Latitude:  45.4641,
			Longitude: 9.1919,
		},
		TimeZone: "Europe/Rome",
		Metadata: model.Metadata{
			IsSaved:     false,
			DeleteAfter: 48 * time.Hour,
		},
	}

	s.mockRepo.On("Save", newDestination).Return(nil)

	createdDestination, err := s.destinationService.CreateDestination(newDestination)

	s.NoError(err)
	s.NotNil(createdDestination)
	s.Equal("Milan Cathedral", createdDestination.Name)
	s.Equal("Piazza Duomo, Milano, Italy", createdDestination.FullAddress)
	s.Equal("Europe/Rome", createdDestination.TimeZone)
	s.Equal(uint(1), createdDestination.UserID)
}

func (s *DestinationServiceTestSuite) Test_should_return_error_when_save_fails() {
	newDestination := &model.Destination{
		PublicID:    uuid.New(),
		UserID:      1,
		FullAddress: "Test Address",
		Name:        "Test Place",
		TimeZone:    "Europe/Rome",
	}

	s.mockRepo.On("Save", mock.Anything).Return(errors.New("database error"))

	createdDestination, err := s.destinationService.CreateDestination(newDestination)

	s.Error(err)
	s.Nil(createdDestination)
	s.Contains(err.Error(), "database")
}

func (s *DestinationServiceTestSuite) Test_should_return_error_when_duplicate_google_place_id() {
	newDestination := &model.Destination{
		PublicID:    uuid.New(),
		UserID:      1,
		FullAddress: "Via Roma 123, Milano, Italy",
		Name:        "Office",
		AddressComponents: model.AddressComponents{
			FormattedAddress: "Via Roma 123, 20100 Milano MI, Italy",
			GooglePlaceId:    "ChIJAWaem...",
		},
		TimeZone: "Europe/Rome",
	}

	s.mockRepo.On("Save", mock.Anything).
		Return(errors.New("duplicate key value violates unique constraint"))

	createdDestination, err := s.destinationService.CreateDestination(newDestination)

	s.Error(err)
	s.Contains(err.Error(), "duplicate")
	s.Nil(createdDestination)
}

func (s *DestinationServiceTestSuite) Test_should_delete_destination_successfully() {
	destinationID := uint(123)

	s.mockRepo.On("Delete", destinationID).Return(nil)

	err := s.destinationService.DeleteDestination(destinationID)

	s.NoError(err)
}

func (s *DestinationServiceTestSuite) Test_should_return_error_when_delete_fails() {
	destinationID := uint(456)

	s.mockRepo.On("Delete", destinationID).Return(errors.New("database error"))

	err := s.destinationService.DeleteDestination(destinationID)

	s.Error(err)
	s.Contains(err.Error(), "database")
}

func (s *DestinationServiceTestSuite) Test_should_return_error_when_destination_not_found_on_delete() {
	destinationID := uint(999)

	s.mockRepo.On("Delete", destinationID).Return(errors.New("record not found"))

	err := s.destinationService.DeleteDestination(destinationID)

	s.Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *DestinationServiceTestSuite) Test_should_not_delete_destination_with_zero_id() {
	destinationID := uint(0)

	s.mockRepo.On("Delete", destinationID).Return(errors.New("invalid id"))

	err := s.destinationService.DeleteDestination(destinationID)

	s.Error(err)
	s.Contains(err.Error(), "invalid")
}

func TestDestinationServiceTestSuite(t *testing.T) {
	suite.Run(t, new(DestinationServiceTestSuite))
}
