package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	_interface "github.com/iacopoGhilardi/amILate/internal/service/interface"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/labstack/echo/v4"
)

type AppointmentHandler struct {
	service _interface.AppointmentServiceInterface
}

func NewAppointmentHandler(service _interface.AppointmentServiceInterface) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

func (h *AppointmentHandler) GetAllAppointments(c echo.Context) error {
	logger.Info("Getting all appointments")
	appointments, err := h.service.GetAllAppointments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	var appointmentDtos []dto.AppointmentDto
	for _, a := range appointments {
		appointmentDtos = append(appointmentDtos, *mapper.MapAppointmentToDto(a))
	}
	return c.JSON(http.StatusOK, commons.Success(appointmentDtos))
}

func (h *AppointmentHandler) GetAppointmentByPublicId(c echo.Context) error {
	id := c.Param("id")
	logger.Info("Getting appointment by id: " + id)
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail("invalid id"))
	}
	appointment, err := h.service.GetAppointmentByPublicId(parsedUuid)
	if err != nil {
		return c.JSON(http.StatusNotFound, commons.Fail("appointment not found"))
	}
	if appointment == nil {
		logger.Info("Appointment not found for id: " + id)
		return c.JSON(http.StatusNotFound, commons.Fail("appointment not found"))
	}
	logger.Info("Appointment found for id: " + id)
	return c.JSON(http.StatusOK, commons.Success(mapper.MapAppointmentToDto(*appointment)))
}

func (h *AppointmentHandler) CreateAppointment(c echo.Context) error {
	logger.Info("Creating appointment")
	var createDto dto.CreateAppointmentRequestDto
	if err := c.Bind(&createDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	createdApp, err := h.service.CreateAppointment(mapper.MapFromCreateAppointmentRequest(createDto))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	logger.Info("Appointment created: " + createdApp.PublicId.String())
	return c.JSON(http.StatusCreated, commons.Success(mapper.MapAppointmentToDto(*createdApp)))
}

func (h *AppointmentHandler) DeleteAppointment(c echo.Context) error {
	id := c.Param("id")
	logger.Info("Deleting appointment by id: " + id)
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail("invalid id"))
	}
	if err := h.service.DeleteAppointmentFromPublicId(parsedUuid); err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}
