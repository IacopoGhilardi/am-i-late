package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/test/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	mockUserService *mocks.UserServiceInterface
	mockAuthService *mocks.AuthServiceInterface
	handler         *handler.UserHandler
	echo            *echo.Echo
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.mockUserService = new(mocks.UserServiceInterface)
	s.mockAuthService = new(mocks.AuthServiceInterface)

	s.handler = handler.NewUserHandler(s.mockUserService, nil)
	s.echo = echo.New()
}

func (s *UserHandlerTestSuite) TearDownTest() {
	s.mockUserService.AssertExpectations(s.T())
}

func (s *UserHandlerTestSuite) createContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
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

func (s *UserHandlerTestSuite) createSampleUser() *model.User {
	return &model.User{
		BaseModel: model.BaseModel{ID: 1},
		Email:     "test@example.com",
		Password:  "hashedpassword123",
		Name:      "Test User",
	}
}

func (s *UserHandlerTestSuite) createSampleUserWithID(id uint, email string) *model.User {
	return &model.User{
		BaseModel: model.BaseModel{ID: id},
		Email:     email,
		Password:  "hashedpassword123",
		Name:      "User " + email,
	}
}

func (s *UserHandlerTestSuite) Test_should_register_user_successfully() {
	requestBody := `{
		"email": "newuser@example.com",
		"password": "Password123!",
		"confirm_password": "Password123!",
		"name": "New User",
		"age_confirmed": true,
		"privacy_accepted": true,
		"terms_accepted": true,
		"location_permission": true
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/register", requestBody)
	err := s.handler.Register(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_invalid_json_on_register() {
	invalidJSON := `{"invalid json`

	c, responseRecorder := s.createContext(http.MethodPost, "/auth/register", invalidJSON)
	err := s.handler.Register(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)

	var errorResponse handler.ErrorResponse
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &errorResponse)
	s.NoError(err)
	s.Equal("Invalid request body", errorResponse.Error)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_missing_required_fields_on_register() {
	requestBody := `{
		"email": "newuser@example.com"
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/register", requestBody)
	err := s.handler.Register(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_register_with_all_permissions_accepted() {
	requestBody := `{
		"email": "newuser@example.com",
		"password": "Password123!",
		"confirm_password": "Password123!",
		"name": "New User",
		"age_confirmed": true,
		"privacy_accepted": true,
		"terms_accepted": true,
		"location_permission": true
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/register", requestBody)
	err := s.handler.Register(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_register_with_location_permission_false() {
	requestBody := `{
		"email": "newuser@example.com",
		"password": "Password123!",
		"confirm_password": "Password123!",
		"name": "New User",
		"age_confirmed": true,
		"privacy_accepted": true,
		"terms_accepted": true,
		"location_permission": false
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/register", requestBody)
	err := s.handler.Register(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_login_user_successfully() {
	requestBody := `{
		"email": "test@example.com",
		"password": "Password123!"
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/login", requestBody)
	err := s.handler.Login(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_invalid_json_on_login() {
	invalidJSON := `{"invalid json`

	c, responseRecorder := s.createContext(http.MethodPost, "/auth/login", invalidJSON)
	err := s.handler.Login(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)

	var errorResponse handler.ErrorResponse
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &errorResponse)
	s.NoError(err)
	s.Equal("Invalid request body", errorResponse.Error)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_missing_email_on_login() {
	requestBody := `{
		"password": "Password123!"
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/login", requestBody)
	err := s.handler.Login(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_missing_password_on_login() {
	requestBody := `{
		"email": "test@example.com"
	}`

	c, _ := s.createContext(http.MethodPost, "/auth/login", requestBody)
	err := s.handler.Login(c)

	s.NoError(err)
}

func (s *UserHandlerTestSuite) Test_should_get_all_users_successfully() {
	user1 := s.createSampleUserWithID(1, "user1@example.com")
	user2 := s.createSampleUserWithID(2, "user2@example.com")
	user3 := s.createSampleUserWithID(3, "user3@example.com")

	users := []model.User{*user1, *user2, *user3}

	s.mockUserService.On("GetAllUsers").Return(users, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/users", "")
	err := s.handler.GetAllUsers(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result []dto.UserDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Len(result, 3)
	s.Equal("user1@example.com", result[0].Email)
	s.Equal("user2@example.com", result[1].Email)
	s.Equal("user3@example.com", result[2].Email)
}

func (s *UserHandlerTestSuite) Test_should_return_empty_array_when_no_users() {
	s.mockUserService.On("GetAllUsers").Return([]model.User{}, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/users", "")
	err := s.handler.GetAllUsers(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result []dto.UserDto
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Empty(result)
}

func (s *UserHandlerTestSuite) Test_should_return_error_when_get_all_users_fails() {
	s.mockUserService.On("GetAllUsers").Return(nil, errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodGet, "/users", "")
	err := s.handler.GetAllUsers(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *UserHandlerTestSuite) Test_should_get_user_by_id_successfully() {
	user := s.createSampleUser()

	s.mockUserService.On("GetUserByID", uint(1)).Return(user, nil)

	c, responseRecorder := s.createContext(http.MethodGet, "/users/1", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := s.handler.GetUserByID(c)

	s.NoError(err)
	s.Equal(http.StatusOK, responseRecorder.Code)

	var result model.User
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Equal("test@example.com", result.Email)
	s.Equal("Test User", result.Name)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_invalid_id() {
	c, responseRecorder := s.createContext(http.MethodGet, "/users/invalid", "")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := s.handler.GetUserByID(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "invalid id")
}

func (s *UserHandlerTestSuite) Test_should_return_not_found_when_user_does_not_exist() {
	s.mockUserService.On("GetUserByID", uint(999)).Return(nil, errors.New("not found"))

	c, responseRecorder := s.createContext(http.MethodGet, "/users/999", "")
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := s.handler.GetUserByID(c)

	s.NoError(err)
	s.Equal(http.StatusNotFound, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "user not found")
}

func (s *UserHandlerTestSuite) Test_should_create_user_successfully() {
	requestBody := `{
		"email": "newuser@example.com",
		"password": "password123",
		"name": "New User"
	}`

	createdUser := s.createSampleUser()
	createdUser.Email = "newuser@example.com"
	createdUser.Name = "New User"

	s.mockUserService.On("CreateUser", mock.AnythingOfType("*model.User")).Return(createdUser, nil)

	c, responseRecorder := s.createContext(http.MethodPost, "/users", requestBody)
	err := s.handler.CreateUser(c)

	s.NoError(err)
	s.Equal(http.StatusCreated, responseRecorder.Code)

	var result model.User
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	s.NoError(err)
	s.Equal("newuser@example.com", result.Email)
	s.Equal("New User", result.Name)
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_invalid_json_on_create() {
	invalidJSON := `{"invalid json`

	c, responseRecorder := s.createContext(http.MethodPost, "/users", invalidJSON)
	err := s.handler.CreateUser(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
}

func (s *UserHandlerTestSuite) Test_should_return_error_when_create_user_fails() {
	requestBody := `{
		"email": "newuser@example.com",
		"password": "password123",
		"name": "New User"
	}`

	s.mockUserService.On("CreateUser", mock.AnythingOfType("*model.User")).Return(nil, errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodPost, "/users", requestBody)
	err := s.handler.CreateUser(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *UserHandlerTestSuite) Test_should_delete_user_successfully() {
	s.mockUserService.On("DeleteUser", uint(1)).Return(nil)

	c, responseRecorder := s.createContext(http.MethodDelete, "/users/1", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := s.handler.DeleteUser(c)

	s.NoError(err)
	s.Equal(http.StatusNoContent, responseRecorder.Code)
	s.Empty(responseRecorder.Body.String())
}

func (s *UserHandlerTestSuite) Test_should_return_bad_request_for_invalid_id_on_delete() {
	c, responseRecorder := s.createContext(http.MethodDelete, "/users/invalid", "")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := s.handler.DeleteUser(c)

	s.NoError(err)
	s.Equal(http.StatusBadRequest, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "invalid id")
}

func (s *UserHandlerTestSuite) Test_should_return_error_when_delete_user_fails() {
	s.mockUserService.On("DeleteUser", uint(1)).Return(errors.New("database error"))

	c, responseRecorder := s.createContext(http.MethodDelete, "/users/1", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := s.handler.DeleteUser(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "database error")
}

func (s *UserHandlerTestSuite) Test_should_return_error_when_user_not_found_on_delete() {
	s.mockUserService.On("DeleteUser", uint(999)).Return(errors.New("user not found"))

	c, responseRecorder := s.createContext(http.MethodDelete, "/users/999", "")
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := s.handler.DeleteUser(c)

	s.NoError(err)
	s.Equal(http.StatusInternalServerError, responseRecorder.Code)
	s.Contains(responseRecorder.Body.String(), "user not found")
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
