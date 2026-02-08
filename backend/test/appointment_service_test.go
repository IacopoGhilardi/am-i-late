package test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/iacopoGhilardi/amILate/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AppointmentServiceTestSuite struct {
	suite.Suite
	mockRepo           *mocks.AppointmentRepositoryInterface
	appointmentService *service.AppointmentService
}

func (s *AppointmentServiceTestSuite) SetupTest() {
	s.mockRepo = new(mocks.AppointmentRepositoryInterface)
	s.appointmentService = service.NewAppointmentService(s.mockRepo)
}

func (s *AppointmentServiceTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func (s *AppointmentServiceTestSuite) Test_should_get_all_appointments_successfully() {
	expectedAppointments := []model.Appointment{
		{
			PublicId:               uuid.New(),
			UserID:                 1,
			DestinationID:          10,
			ScheduledAt:            "2026-02-10T09:00:00Z",
			TransportMode:          "car",
			EstimatedTravelMinutes: 30,
			EstimatedTravelRange:   5,
			NotificationState:      "pending",
			Status:                 "scheduled",
			DeleteAfter:            "2026-02-10T18:00:00Z",
			GeoFenceId:             "fence_123",
		},
		{
			PublicId:               uuid.New(),
			UserID:                 1,
			DestinationID:          20,
			ScheduledAt:            "2026-02-11T14:00:00Z",
			TransportMode:          "public_transport",
			EstimatedTravelMinutes: 45,
			EstimatedTravelRange:   10,
			NotificationState:      "monitoring",
			Status:                 "scheduled",
			DeleteAfter:            "2026-02-11T20:00:00Z",
			GeoFenceId:             "fence_456",
		},
	}

	s.mockRepo.On("FindAll").Return(expectedAppointments, nil)

	appointments, err := s.appointmentService.GetAllAppointments()

	s.NoError(err)
	s.NotNil(appointments)
	s.Len(appointments, 2)
	s.Equal("car", appointments[0].TransportMode)
	s.Equal("public_transport", appointments[1].TransportMode)
}

func (s *AppointmentServiceTestSuite) Test_should_return_empty_list_when_no_appointments() {
	s.mockRepo.On("FindAll").Return([]model.Appointment{}, nil)

	appointments, err := s.appointmentService.GetAllAppointments()

	s.NoError(err)
	s.NotNil(appointments)
	s.Empty(appointments)
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_findall_fails() {
	s.mockRepo.On("FindAll").Return(nil, errors.New("database error"))

	appointments, err := s.appointmentService.GetAllAppointments()

	s.Error(err)
	s.Nil(appointments)
	s.Contains(err.Error(), "database")
}

func (s *AppointmentServiceTestSuite) Test_should_find_appointment_by_id() {
	appointmentID := uint(123)
	publicID := uuid.New()

	expectedAppointment := &model.Appointment{
		BaseModel: model.BaseModel{
			ID: appointmentID,
		},
		PublicId:               publicID,
		UserID:                 1,
		DestinationID:          10,
		ScheduledAt:            "2026-02-10T09:00:00Z",
		TransportMode:          "car",
		EstimatedTravelMinutes: 30,
		EstimatedTravelRange:   5,
		NotificationState:      "pending",
		Status:                 "scheduled",
		DeleteAfter:            "2026-02-10T18:00:00Z",
		GeoFenceId:             "fence_123",
	}

	s.mockRepo.On("Find", appointmentID).Return(expectedAppointment, nil)

	appointment, err := s.appointmentService.GetAppointmentByID(appointmentID)

	s.NoError(err)
	s.NotNil(appointment)
	s.Equal(appointmentID, appointment.ID)
	s.Equal(publicID, appointment.PublicId)
	s.Equal("car", appointment.TransportMode)
	s.Equal("scheduled", appointment.Status)
}

func (s *AppointmentServiceTestSuite) Test_should_return_nil_when_appointment_not_found_by_id() {
	appointmentID := uint(999)

	s.mockRepo.On("Find", appointmentID).Return(nil, nil)

	appointment, err := s.appointmentService.GetAppointmentByID(appointmentID)

	s.NoError(err)
	s.Nil(appointment)
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_find_by_id_fails() {
	appointmentID := uint(123)

	s.mockRepo.On("Find", appointmentID).Return(nil, errors.New("database error"))

	appointment, err := s.appointmentService.GetAppointmentByID(appointmentID)

	s.Error(err)
	s.Nil(appointment)
	s.Contains(err.Error(), "database")
}

func (s *AppointmentServiceTestSuite) Test_should_find_appointment_by_public_id() {
	publicID := uuid.New()

	expectedAppointment := &model.Appointment{
		PublicId:               publicID,
		UserID:                 1,
		DestinationID:          10,
		ScheduledAt:            "2026-02-10T09:00:00Z",
		TransportMode:          "foot",
		EstimatedTravelMinutes: 15,
		EstimatedTravelRange:   3,
		NotificationState:      "pending",
		Status:                 "scheduled",
		DeleteAfter:            "2026-02-10T18:00:00Z",
		GeoFenceId:             "fence_789",
	}

	s.mockRepo.On("FindByPublicId", publicID).Return(expectedAppointment, nil)

	appointment, err := s.appointmentService.GetAppointmentByPublicId(publicID)

	s.NoError(err)
	s.NotNil(appointment)
	s.Equal(publicID, appointment.PublicId)
	s.Equal("foot", appointment.TransportMode)
}

func (s *AppointmentServiceTestSuite) Test_should_return_nil_when_appointment_not_found_by_public_id() {
	publicID := uuid.New()

	s.mockRepo.On("FindByPublicId", publicID).Return(nil, nil)

	appointment, err := s.appointmentService.GetAppointmentByPublicId(publicID)

	s.NoError(err)
	s.Nil(appointment)
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_find_by_public_id_fails() {
	publicID := uuid.New()

	s.mockRepo.On("FindByPublicId", publicID).Return(nil, errors.New("database error"))

	appointment, err := s.appointmentService.GetAppointmentByPublicId(publicID)

	s.Error(err)
	s.Nil(appointment)
	s.Contains(err.Error(), "database")
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_public_id_is_invalid() {
	invalidID := uuid.Nil

	s.mockRepo.On("FindByPublicId", invalidID).Return(nil, errors.New("invalid uuid"))

	appointment, err := s.appointmentService.GetAppointmentByPublicId(invalidID)

	s.Error(err)
	s.Nil(appointment)
	s.Contains(err.Error(), "invalid")
}

func (s *AppointmentServiceTestSuite) Test_should_create_appointment_successfully() {
	newAppointment := &model.Appointment{
		PublicId:               uuid.New(),
		UserID:                 1,
		DestinationID:          10,
		ScheduledAt:            "2026-02-15T10:00:00Z",
		TransportMode:          "car",
		EstimatedTravelMinutes: 20,
		EstimatedTravelRange:   5,
		LastTravelUpdateAt:     "2026-02-08T12:00:00Z",
		NotificationState:      "pending",
		Status:                 "scheduled",
		DeleteAfter:            "2026-02-15T20:00:00Z",
		GeoFenceId:             "fence_new",
	}

	s.mockRepo.On("Save", newAppointment).Return(nil)

	createdAppointment, err := s.appointmentService.CreateAppointment(newAppointment)

	s.NoError(err)
	s.NotNil(createdAppointment)
	s.Equal(uint(1), createdAppointment.UserID)
	s.Equal(uint(10), createdAppointment.DestinationID)
	s.Equal("car", createdAppointment.TransportMode)
	s.Equal("scheduled", createdAppointment.Status)
	s.Equal("pending", createdAppointment.NotificationState)
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_create_fails() {
	newAppointment := &model.Appointment{
		PublicId:          uuid.New(),
		UserID:            1,
		DestinationID:     10,
		ScheduledAt:       "2026-02-15T10:00:00Z",
		TransportMode:     "public_transport",
		NotificationState: "pending",
		Status:            "scheduled",
	}

	s.mockRepo.On("Save", mock.Anything).Return(errors.New("database error"))

	createdAppointment, err := s.appointmentService.CreateAppointment(newAppointment)

	s.Error(err)
	s.NotNil(createdAppointment)
	s.Contains(err.Error(), "database")
}

func (s *AppointmentServiceTestSuite) Test_should_validate_transport_mode() {
	appointment := model.Appointment{}

	s.True(appointment.ValidateTransport("car"))
	s.True(appointment.ValidateTransport("public_transport"))
	s.True(appointment.ValidateTransport("foot"))
	s.False(appointment.ValidateTransport("airplane"))
	s.False(appointment.ValidateTransport("invalid"))
}

func (s *AppointmentServiceTestSuite) Test_should_validate_status() {
	appointment := model.Appointment{}

	s.True(appointment.ValidateStatus("scheduled"))
	s.True(appointment.ValidateStatus("completed"))
	s.True(appointment.ValidateStatus("cancelled"))
	s.False(appointment.ValidateStatus("pending"))
	s.False(appointment.ValidateStatus("invalid"))
}

func (s *AppointmentServiceTestSuite) Test_should_validate_notification_state() {
	appointment := model.Appointment{}

	s.True(appointment.ValidateNotificationState("pending"))
	s.True(appointment.ValidateNotificationState("monitoring"))
	s.True(appointment.ValidateNotificationState("sent"))
	s.True(appointment.ValidateNotificationState("cancelled"))
	s.False(appointment.ValidateNotificationState("invalid"))
}

func (s *AppointmentServiceTestSuite) Test_should_delete_appointment_successfully() {
	appointmentID := uint(123)

	s.mockRepo.On("Delete", appointmentID).Return(nil)

	err := s.appointmentService.DeleteAppointment(appointmentID)

	s.NoError(err)
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_delete_fails() {
	appointmentID := uint(456)

	s.mockRepo.On("Delete", appointmentID).Return(errors.New("database error"))

	err := s.appointmentService.DeleteAppointment(appointmentID)

	s.Error(err)
	s.Contains(err.Error(), "database")
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_appointment_not_found_on_delete() {
	appointmentID := uint(999)

	s.mockRepo.On("Delete", appointmentID).Return(errors.New("record not found"))

	err := s.appointmentService.DeleteAppointment(appointmentID)

	s.Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *AppointmentServiceTestSuite) Test_should_delete_appointment_by_public_id_successfully() {
	publicID := uuid.New()
	appointmentID := uint(123)

	appointment := &model.Appointment{
		BaseModel: model.BaseModel{ID: appointmentID},
		PublicId:  publicID,
		UserID:    1,
		Status:    "scheduled",
	}

	s.mockRepo.On("FindByPublicId", publicID).Return(appointment, nil)
	s.mockRepo.On("Delete", appointmentID).Return(nil)

	err := s.appointmentService.DeleteAppointmentFromPublicId(publicID)

	s.NoError(err)
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_appointment_not_found_by_public_id_on_delete() {
	publicID := uuid.New()

	s.mockRepo.On("FindByPublicId", publicID).Return(nil, errors.New("record not found"))

	err := s.appointmentService.DeleteAppointmentFromPublicId(publicID)

	s.Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *AppointmentServiceTestSuite) Test_should_return_error_when_delete_by_public_id_fails() {
	publicID := uuid.New()
	appointmentID := uint(123)

	appointment := &model.Appointment{
		BaseModel: model.BaseModel{ID: appointmentID},
		PublicId:  publicID,
	}

	s.mockRepo.On("FindByPublicId", publicID).Return(appointment, nil)
	s.mockRepo.On("Delete", appointmentID).Return(errors.New("database error"))

	err := s.appointmentService.DeleteAppointmentFromPublicId(publicID)

	s.Error(err)
	s.Contains(err.Error(), "database")
}

func TestAppointmentServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AppointmentServiceTestSuite))
}
