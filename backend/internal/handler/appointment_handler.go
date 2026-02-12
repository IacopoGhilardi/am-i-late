package handler

import (
	"net/http"

	"github.com/google/uuid"
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
	var appointmentDtos []dto.AppointmentDto
	for _, a := range appointments {
		appointmentDtos = append(appointmentDtos, *mapper.MapAppointmentToDto(a))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, appointmentDtos)
}

func (h *AppointmentHandler) GetAppointmentByPublicId(c echo.Context) error {
	id := c.Param("id")
	parsedUuid, err := uuid.Parse(id)
	logger.Info("Getting appointment by id: " + id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	appointment, err := h.service.GetAppointmentByPublicId(uuid.Must(parsedUuid, nil))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "appointment not found"})
	}
	if appointment == nil {
		logger.Info("Appointment not found for id: " + id + "")
		return c.JSON(http.StatusNotFound, map[string]string{"error": "appointment not found"})
	}
	logger.Info("Appointment found for id: " + id + "")
	appointmentDto := mapper.MapAppointmentToDto(*appointment)
	return c.JSON(http.StatusOK, appointmentDto)
}

func (h *AppointmentHandler) CreateAppointment(c echo.Context) error {
	var dto dto.CreateAppointmentRequestDto
	logger.Info("Creating appointment")
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	createdApp, err := h.service.CreateAppointment(mapper.MapFromCreateAppointmentRequest(dto))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	logger.Info("Appointment created: " + createdApp.PublicId.String())
	appointmentDto := mapper.MapAppointmentToDto(*createdApp)
	return c.JSON(http.StatusCreated, appointmentDto)
}

func (h *AppointmentHandler) DeleteAppointment(c echo.Context) error {
	id := c.Param("id")
	parsedUuid, err := uuid.Parse(id)
	logger.Info("Deleting appointment by id: " + id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := h.service.DeleteAppointmentFromPublicId(parsedUuid); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
