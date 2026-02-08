package handler

import (
	"net/http"
	"strings"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/model"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/iacopoGhilardi/amILate/internal/utils"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service     *service.UserService
	authService *service.AuthService
}

func NewUserHandler(service *service.UserService, authService *service.AuthService) *UserHandler {
	return &UserHandler{
		service:     service,
		authService: authService,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	var registerDto dto.RegistrationDto

	if err := c.Bind(&registerDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	if err := c.Validate(&registerDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
	}

	response, err := h.authService.Register(registerDto)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			return c.JSON(http.StatusConflict, ErrorResponse{
				Error:   "Email already registered",
				Message: err.Error(),
			})
		}
		if strings.Contains(err.Error(), "do not match") {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Passwords do not match",
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Registration failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c echo.Context) error {
	var loginDto dto.LoginDto

	// Bind e validazione
	if err := c.Bind(&loginDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	if err := c.Validate(&loginDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
	}

	response, err := h.authService.Login(loginDto)
	if err != nil {
		if strings.Contains(err.Error(), "invalid email or password") {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error:   "Invalid credentials",
				Message: "Invalid email or password",
			})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Login failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := utils.ParseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdUser, err := h.service.CreateUser(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := utils.ParseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.service.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
