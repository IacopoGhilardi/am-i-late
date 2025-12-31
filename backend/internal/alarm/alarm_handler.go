package alarm

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AlarmHandler struct {
	service *AlarmService
}

func NewAlarmHandler(service *AlarmService) *AlarmHandler {
	return &AlarmHandler{service: service}
}

func (h *AlarmHandler) GetAllAlarms(c echo.Context) error {
	alarms, err := h.service.GetAllAlarms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, alarms)
}

func (h *AlarmHandler) GetAlarmByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	alarm, err := h.service.GetAlarmByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "alarm not found"})
	}
	return c.JSON(http.StatusOK, alarm)
}

func (h *AlarmHandler) CreateAlarm(c echo.Context) error {
	var a CreateAlarmRequestDto
	if err := c.Bind(&a); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	created, err := h.service.CreateAlarm(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *AlarmHandler) DeleteAlarm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.service.DeleteAlarm(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
