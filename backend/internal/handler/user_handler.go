package handler

import (
	"net/http"
	"strings"

	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	"github.com/iacopoGhilardi/amILate/internal/model"
	_interface "github.com/iacopoGhilardi/amILate/internal/service/interface"
	"github.com/iacopoGhilardi/amILate/internal/utils"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service     _interface.UserServiceInterface
	authService _interface.AuthServiceInterface
}

func NewUserHandler(service _interface.UserServiceInterface, authService _interface.AuthServiceInterface) *UserHandler {
	return &UserHandler{
		service:     service,
		authService: authService,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	var registerDto dto.RegistrationDto
	if err := c.Bind(&registerDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	if err := c.Validate(&registerDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	response, err := h.authService.Register(registerDto)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			return c.JSON(http.StatusConflict, commons.Fail("email already registered"))
		}
		if strings.Contains(err.Error(), "do not match") {
			return c.JSON(http.StatusBadRequest, commons.Fail("passwords do not match"))
		}
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.JSON(http.StatusCreated, commons.Success(response))
}

func (h *UserHandler) Login(c echo.Context) error {
	var loginDto dto.LoginDto
	if err := c.Bind(&loginDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	if err := c.Validate(&loginDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	response, err := h.authService.Login(loginDto)
	if err != nil {
		if strings.Contains(err.Error(), "invalid email or password") {
			return c.JSON(http.StatusUnauthorized, commons.Fail("invalid email or password"))
		}
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.JSON(http.StatusOK, commons.Success(response))
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	logger.Info("Getting all users")
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	var userDtos []dto.UserDto
	for _, user := range users {
		userDtos = append(userDtos, *mapper.MapUserToDto(user))
	}
	return c.JSON(http.StatusOK, commons.Success(userDtos))
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := utils.ParseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail("invalid id"))
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, commons.Fail("user not found"))
	}
	return c.JSON(http.StatusOK, commons.Success(mapper.MapUserToDto(*user)))
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	createdUser, err := h.service.CreateUser(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.JSON(http.StatusCreated, commons.Success(createdUser))
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := utils.ParseIDParam(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail("invalid id"))
	}
	if err := h.service.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}
