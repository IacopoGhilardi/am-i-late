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

type DestinationHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.DestinationServiceInterface
	handler     *handler.DestinationHandler
	echo        *echo.Echo
}

func (s *DestinationHandlerTestSuite) SetupTest() {
	s.mockService = new(mocks.DestinationServiceInterface)
	s.handler = handler.NewDestinationHandler(s.mockService)
	s.echo = echo.New()
}

func (s *DestinationHandlerTestSuite) TearDownTest() {
	s.mockService.AssertExpectations(s.T())
}

func (s *DestinationHandlerTestSuite) createContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
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

func (s *DestinationHandlerTestSuite) createSampleDestination() model.Destination {
	return model.Destination{
		BaseModel:   model.BaseModel{ID: 1},
		PublicID:    uuid.New(),
		Name:        "Office",
		FullAddress: "Via Roma 1, Milan",
		Location: model.Location{
			Latitude:  45.4642,
			Longitude: 9.1900,
		},
	}
}

func (s *DestinationHandlerTestSuite) createSampleDestinationWithID(id uint) model.Destination {
	return model.Destination{
		BaseModel:   model.BaseModel{ID: id},
		PublicID:    uuid.New(),
		Name:        "Office",
		FullAddress: "Via Roma 1, Milan",
		Location: model.Location{
			Latitude:  45.4642,
			Longitude: 9.1900,
		},
	}
}

func (s *DestinationHandlerTestSuite) Test_should_get_all_destinations_successfully() {
	destination1 := s.createSampleDestinationWithID(1)
	destination2 := s.createSampleDestinationWithID(2)
	destination2.Name = "Home"
	destination2.FullAddress = "Via Milano 10, Milan"
	destination2.Location.Latitude = 45.4700
	destination2.Location.Longitude = 9.1950

	destinations := []model.Destination{destination1, destination2}

	s.mockService.On("GetAllDestinations").Return(destinations, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/destinations", "")
	err := s.handler.GetAllDestinations(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result []dto.DestinationDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Len(result, 2)
	s.Equal("Office", result[0].Name)
	s.Equal("Home", result[1].Name)
}

func (s *DestinationHandlerTestSuite) Test_should_return_empty_array_when_no_destinations() {
	s.mockService.On("GetAllDestinations").Return([]model.Destination{}, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/destinations", "")
	err := s.handler.GetAllDestinations(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result []dto.DestinationDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Empty(result)
}

func (s *DestinationHandlerTestSuite) Test_should_return_error_when_get_all_destinations_fails() {
	s.mockService.On("GetAllDestinations").Return(nil, errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodGet, "/destinations", "")
	err := s.handler.GetAllDestinations(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *DestinationHandlerTestSuite) Test_should_get_multiple_destinations_with_different_locations() {
	dest1 := s.createSampleDestinationWithID(1)
	dest2 := s.createSampleDestinationWithID(2)
	dest3 := s.createSampleDestinationWithID(3)

	dest1.Name = "Office Milan"
	dest1.Location.Latitude = 45.4642
	dest1.Location.Longitude = 9.1900

	dest2.Name = "Home Rome"
	dest2.Location.Latitude = 41.9028
	dest2.Location.Longitude = 12.4964

	dest3.Name = "Gym Florence"
	dest3.Location.Latitude = 43.7696
	dest3.Location.Longitude = 11.2558

	destinations := []model.Destination{dest1, dest2, dest3}

	s.mockService.On("GetAllDestinations").Return(destinations, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/destinations", "")
	err := s.handler.GetAllDestinations(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result []dto.DestinationDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Len(result, 3)
}

// ==================== Tests for GetDestinationByID ====================

func (s *DestinationHandlerTestSuite) Test_should_get_destination_by_id_successfully() {
	destination := s.createSampleDestinationWithID(1)

	s.mockService.On("GetDestinationByID", uint(1)).Return(&destination, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/destinations/1", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := s.handler.GetDestinationByID(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.NotNil(result)
}

func (s *DestinationHandlerTestSuite) Test_should_return_bad_request_for_invalid_id() {
	c, responseRecorder := s.createContext(http.MethodGet, "/destinations/invalid", "")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := s.handler.GetDestinationByID(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "invalid id")
}

func (s *DestinationHandlerTestSuite) Test_should_return_not_found_when_destination_does_not_exist() {
	s.mockService.On("GetDestinationByID", uint(999)).Return(nil, errors.New("not found"))

	c, responseRecorder := s.createContext(http.MethodGet, "/destinations/999", "")
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := s.handler.GetDestinationByID(c)

	s.NoError(err)
	s.Equal(http.StatusNotFound, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "destination not found")
}

func (s *DestinationHandlerTestSuite) Test_should_create_destination_successfully() {
	requestBody := `{
		"name": "New Office",
		"address": "Via Nuova 1, Milan",
		"latitude": 45.4642,
		"longitude": 9.1900
	}`

	createdDestination := s.createSampleDestination()
	createdDestination.Name = "New Office"

	s.mockService.On("CreateDestination", mock.AnythingOfType("*model.Destination")).Return(&createdDestination, nil)

	c, responseRecorder := s.createContext(http.MethodPost, "/destinations", requestBody)
	err := s.handler.CreateDestination(c)

	s.NoError(err)
	s.Equal(http.StatusCreated, responseRecorder.Code)

	var result map[string]interface{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.NotNil(result)
}

func (s *DestinationHandlerTestSuite) Test_should_return_bad_request_for_invalid_json_on_create() {
	invalidJSON := `{"invalid json`

	c, responseRecorder := s.createContext(http.MethodPost, "/destinations", invalidJSON)
	err := s.handler.CreateDestination(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
}

func (s *DestinationHandlerTestSuite) Test_should_return_error_when_create_destination_fails() {
	requestBody := `{
		"name": "New Office",
		"address": "Via Nuova 1, Milan",
		"latitude": 45.4642,
		"longitude": 9.1900
	}`

	s.mockService.On("CreateDestination", mock.AnythingOfType("*model.Destination")).Return(nil, errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodPost, "/destinations", requestBody)
	err := s.handler.CreateDestination(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *DestinationHandlerTestSuite) Test_should_create_destination_with_various_coordinates() {
	requestBody := `{
		"name": "Remote Location",
		"address": "North Pole",
		"latitude": 90.0,
		"longitude": 0.0
	}`

	createdDestination := s.createSampleDestination()
	createdDestination.Name = "Remote Location"
	createdDestination.Location.Latitude = 90.0
	createdDestination.Location.Longitude = 0.0

	s.mockService.On("CreateDestination", mock.AnythingOfType("*model.Destination")).Return(&createdDestination, nil)

	c, responseRecorder := s.createContext(http.MethodPost, "/destinations", requestBody)
	err := s.handler.CreateDestination(c)

	s.NoError(err)
	s.Equal(http.StatusCreated, responseRecorder.Code)
}

func (s *DestinationHandlerTestSuite) Test_should_delete_destination_successfully() {
	s.mockService.On("DeleteDestination", uint(1)).Return(nil)

	c, responseRecorder := s.createContext(http.MethodDelete, "/destinations/1", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := s.handler.DeleteDestination(c)

	s.NoError(err)
	s.Equal(http.StatusNoContent, responseRecorder.Code)
	s.Empty(responseRecorder.Body.String())
}

func (s *DestinationHandlerTestSuite) Test_should_return_bad_request_for_invalid_id_on_delete() {
	c, responseRecorder := s.createContext(http.MethodDelete, "/destinations/invalid", "")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := s.handler.DeleteDestination(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "invalid id")
}

func (s *DestinationHandlerTestSuite) Test_should_return_error_when_delete_destination_fails() {
	s.mockService.On("DeleteDestination", uint(1)).Return(errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodDelete, "/destinations/1", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := s.handler.DeleteDestination(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *DestinationHandlerTestSuite) Test_should_return_error_when_destination_not_found_on_delete() {
	s.mockService.On("DeleteDestination", uint(999)).Return(errors.New("destination not found"))

	c, responseRecorder := s.createContext(http.MethodDelete, "/destinations/999", "")
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := s.handler.DeleteDestination(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "destination not found")
}

func TestDestinationHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DestinationHandlerTestSuite))
}
