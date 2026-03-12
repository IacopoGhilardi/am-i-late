package handler

import (
	"net/http"
	"strings"

	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	"github.com/iacopoGhilardi/amILate/internal/middleware"
	"github.com/iacopoGhilardi/amILate/internal/model"
	_interface "github.com/iacopoGhilardi/amILate/internal/service/interface"
	"github.com/iacopoGhilardi/amILate/internal/utils"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service         _interface.UserServiceInterface
	authService     _interface.AuthServiceInterface
	locationService _interface.UserLocationServiceInterface
}

func NewUserHandler(
	service _interface.UserServiceInterface,
	authService _interface.AuthServiceInterface,
	locationService _interface.UserLocationServiceInterface) *UserHandler {
	return &UserHandler{
		service:         service,
		authService:     authService,
		locationService: locationService,
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

func (h *UserHandler) ForgotPassword(c echo.Context) error {
	var forgotDto dto.ForgotPasswordDto
	if err := c.Bind(&forgotDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	if err := c.Validate(&forgotDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}

	_ = h.authService.ForgotPassword(forgotDto)
	return c.JSON(http.StatusOK, commons.Success("if the email exists, you will receive a reset link"))
}

func (h *UserHandler) ResetPassword(c echo.Context) error {
	var resetDto dto.ResetPasswordDto
	if err := c.Bind(&resetDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	if err := c.Validate(&resetDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	if err := h.authService.ResetPassword(resetDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	return c.JSON(http.StatusOK, commons.Success("password reset successfully"))
}

func (h *UserHandler) UpdateLocation(c echo.Context) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, commons.Fail("unauthorized"))
	}

	var locationDto dto.UpdateLocationDto
	if err := c.Bind(&locationDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	if err := c.Validate(&locationDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}

	if err := h.locationService.UpdateLocation(claims.UserId, locationDto.Latitude, locationDto.Longitude); err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}

	return c.JSON(http.StatusOK, commons.Success("location updated"))
}
