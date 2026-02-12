package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/test/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AppointmentHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.AppointmentServiceInterface
	handler     *handler.AppointmentHandler
	echo        *echo.Echo
}

func (s *AppointmentHandlerTestSuite) SetupTest() {
	s.mockService = new(mocks.AppointmentServiceInterface)
	s.handler = handler.NewAppointmentHandler(s.mockService)
	s.echo = echo.New()
}

func (s *AppointmentHandlerTestSuite) TearDownTest() {
	s.mockService.AssertExpectations(s.T())
}

func (s *AppointmentHandlerTestSuite) createContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	var req *http.Request
	if body != "" {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	rec := httptest.NewRecorder()
	return s.echo.NewContext(req, rec), rec
}

func (s *AppointmentHandlerTestSuite) Test_should_get_appointment_successfully() {
	appointment := s.createSampleAppointment()

	s.mockService.On("GetAppointmentByPublicId", appointment.PublicId).Return(appointment, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/appointments/"+appointment.PublicId.String(), "")
	c.SetParamNames("id")
	c.SetParamValues(appointment.PublicId.String())

	err := s.handler.GetAppointmentByPublicId(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result dto.AppointmentDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.Equal(appointment.PublicId, result.ID)
	s.Equal("car", result.TransportMode)
	s.Equal(30, result.EstimatedTravelMinutes)
}

func (s *AppointmentHandlerTestSuite) Test_should_return_error_when_appointment_not_found_by_id() {
	randomUuid := uuid.New()
	s.mockService.On("GetAppointmentByPublicId", randomUuid).Return(nil, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/appointments/"+randomUuid.String(), "")
	c.SetParamNames("id")
	c.SetParamValues(randomUuid.String())

	err := s.handler.GetAppointmentByPublicId(c)
	s.NoError(err)
	s.Equal(http.StatusNotFound, responseRecorder.Code)
}

func (s *AppointmentHandlerTestSuite) Test_should_create_appointment_successfully() {
	requestBody := `{
		"destination": {
			"name": "Office",
			"address": "Via Roma 1, Milan",
			"latitude": 45.4642,
			"longitude": 9.1900
		},
		"scheduled_at": "2026-02-15T10:00:00Z",
		"transport_mode": "car"
	}`

	createdAppointment := s.createSampleAppointment()
	s.mockService.On("CreateAppointment", mock.AnythingOfType("*model.Appointment")).Return(createdAppointment, nil)

	c, responseRecorder := s.createContext(http.MethodPost, "/appointments", requestBody)
	err := s.handler.CreateAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusCreated, responseRecorder.Code)

	var result dto.AppointmentDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)

	s.NotEqual(uuid.Nil, result.ID)
	s.Equal("car", result.TransportMode)
	s.Equal("scheduled", result.Status)
	s.Equal(30, result.EstimatedTravelMinutes)
}

func (s *AppointmentHandlerTestSuite) Test_should_return_error_when_create_fails() {
	requestBody := `{
		"destination": {
			"name": "Office",
			"address": "Via Roma 1, Milan",
			"latitude": 45.4642,
			"longitude": 9.1900
		},
		"scheduled_at": "2026-02-15T10:00:00Z",
		"transport_mode": "car"
	}`

	s.mockService.On("CreateAppointment", mock.AnythingOfType("*model.Appointment")).Return(nil, errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodPost, "/appointments", requestBody)

	err := s.handler.CreateAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *AppointmentHandlerTestSuite) Test_should_return_bad_request_for_invalid_json() {
	invalidJSON := `{"invalid json`

	c, responseRecorder := s.createContext(http.MethodPost, "/appointments", invalidJSON)
	err := s.handler.CreateAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
}

func (s *AppointmentHandlerTestSuite) Test_should_create_appointment_with_public_transport() {
	requestBody := `{
		"destination": {
			"name": "Station",
			"address": "Via Milano 10",
			"latitude": 45.4652,
			"longitude": 9.1910
		},
		"scheduled_at": "2026-02-16T08:00:00Z",
		"transport_mode": "public_transport"
	}`

	createdAppointment := s.createSampleAppointment()
	createdAppointment.TransportMode = "public_transport"

	s.mockService.On("CreateAppointment", mock.AnythingOfType("*model.Appointment")).Return(createdAppointment, nil)

	c, responseRecorder := s.createContext(http.MethodPost, "/appointments", requestBody)

	err := s.handler.CreateAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusCreated, responseRecorder.Code)

	var result dto.AppointmentDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Equal("public_transport", result.TransportMode)
}

func (s *AppointmentHandlerTestSuite) Test_should_create_appointment_with_foot() {
	requestBody := `{
		"destination": {
			"name": "Park",
			"address": "Via Parco 5",
			"latitude": 45.4700,
			"longitude": 9.1950
		},
		"scheduled_at": "2026-02-17T14:00:00Z",
		"transport_mode": "foot"
	}`

	createdAppointment := s.createSampleAppointment()
	createdAppointment.TransportMode = "foot"

	s.mockService.On("CreateAppointment", mock.AnythingOfType("*model.Appointment")).Return(createdAppointment, nil)

	c, responseRecorder := s.createContext(http.MethodPost, "/appointments", requestBody)

	err := s.handler.CreateAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusCreated, responseRecorder.Code)

	var result dto.AppointmentDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Equal("foot", result.TransportMode)
}

func (s *AppointmentHandlerTestSuite) createSampleAppointment() *model.Appointment {
	return &model.Appointment{
		BaseModel: model.BaseModel{
			ID: 1,
		},
		PublicId:               uuid.New(),
		UserID:                 1,
		DestinationID:          10,
		ScheduledAt:            "2026-02-10T09:00:00Z",
		TransportMode:          "car",
		EstimatedTravelMinutes: 30,
		EstimatedTravelRange:   5,
		LastTravelUpdateAt:     "2026-02-08T12:00:00Z",
		NotificationState:      "pending",
		Status:                 "scheduled",
		DeleteAfter:            "2026-02-10T18:00:00Z",
		GeoFenceId:             "fence_123",
		Destination: &model.Destination{
			BaseModel:   model.BaseModel{ID: 10},
			PublicID:    uuid.New(),
			Name:        "Office",
			FullAddress: "Via Roma 1, Milan",
			Location: model.Location{
				Latitude:  45.4642,
				Longitude: 9.1900,
			},
		},
	}
}

func (s *AppointmentHandlerTestSuite) Test_should_delete_appointment_successfully() {
	publicID := uuid.New()

	s.mockService.On("DeleteAppointmentFromPublicId", publicID).Return(nil)

	c, responseRecorder := s.createContext(http.MethodDelete, "/appointments/"+publicID.String(), "")
	c.SetParamNames("id")
	c.SetParamValues(publicID.String())

	err := s.handler.DeleteAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusNoContent, responseRecorder.Code)
	s.Empty(responseRecorder.Body.String())
}

func (s *AppointmentHandlerTestSuite) Test_should_return_error_for_invalid_uuid_on_delete() {
	c, responseRecorder := s.createContext(http.MethodDelete, "/appointments/invalid-uuid", "")
	c.SetParamNames("id")
	c.SetParamValues("invalid-uuid")

	err := s.handler.DeleteAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "invalid id")
}

func (s *AppointmentHandlerTestSuite) Test_should_return_error_when_delete_fails() {
	publicID := uuid.New()

	s.mockService.On("DeleteAppointmentFromPublicId", publicID).Return(errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodDelete, "/appointments/"+publicID.String(), "")
	c.SetParamNames("id")
	c.SetParamValues(publicID.String())

	err := s.handler.DeleteAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *AppointmentHandlerTestSuite) Test_should_return_error_when_appointment_not_found_on_delete() {
	publicID := uuid.New()

	s.mockService.On("DeleteAppointmentFromPublicId", publicID).Return(errors.New("appointment not found"))

	c, responseRecorder := s.createContext(http.MethodDelete, "/appointments/"+publicID.String(), "")
	c.SetParamNames("id")
	c.SetParamValues(publicID.String())

	err := s.handler.DeleteAppointment(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "appointment not found")
}

func (s *AppointmentHandlerTestSuite) Test_should_get_all_appointments_successfully() {
	appointment1 := s.createSampleAppointment()

	appointment2 := s.createSampleAppointment()
	appointment2.TransportMode = "public_transport"
	appointment2.EstimatedTravelMinutes = 45
	appointment2.Destination.Name = "Home"

	appointments := []model.Appointment{*appointment1, *appointment2}
	s.mockService.On("GetAllAppointments").Return(appointments, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/appointments", "")
	err := s.handler.GetAllAppointments(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result []dto.AppointmentDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Len(result, 2)
	s.Equal("car", result[0].TransportMode)
	s.Equal("public_transport", result[1].TransportMode)
	s.Equal("scheduled", result[0].Status)
	s.Equal(30, result[0].EstimatedTravelMinutes)
	s.Equal(45, result[1].EstimatedTravelMinutes)
}

func (s *AppointmentHandlerTestSuite) createSampleAppointmentWithoutDestination() *model.Appointment {
	return &model.Appointment{
		BaseModel: model.BaseModel{
			ID: 1,
		},
		PublicId:               uuid.New(),
		UserID:                 1,
		DestinationID:          10,
		ScheduledAt:            "2026-02-10T09:00:00Z",
		TransportMode:          "car",
		EstimatedTravelMinutes: 30,
		EstimatedTravelRange:   5,
		LastTravelUpdateAt:     "2026-02-08T12:00:00Z",
		NotificationState:      "pending",
		Status:                 "scheduled",
		DeleteAfter:            "2026-02-10T18:00:00Z",
		GeoFenceId:             "fence_123",
		Destination:            nil,
	}
}

func (s *AppointmentHandlerTestSuite) createSampleDestination() *model.Destination {
	return &model.Destination{
		BaseModel:   model.BaseModel{ID: 10},
		PublicID:    uuid.New(),
		Name:        "Office",
		FullAddress: "Via Roma 1, Milan",
		Location: model.Location{
			Latitude:  45.4642,
			Longitude: 9.1900,
		},
	}
}

func TestAppointmentHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AppointmentHandlerTestSuite))
}
