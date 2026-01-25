package handler

import (
	"net/http"
	"strconv"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	service2 "github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/iacopoGhilardi/amILate/pkg/logger"
	"github.com/labstack/echo/v4"
)

type DestinationHandler struct {
	service *service2.DestinationService
}

func NewDestinationHandler(service *service2.DestinationService) *DestinationHandler {
	return &DestinationHandler{service: service}
}

func (h *DestinationHandler) GetAllDestinations(c echo.Context) error {
	logger.Info("Getting all destinations")
	dests, err := h.service.GetAllDestinations()
	var destinationDtos []dto.DestinationDto
	for _, d := range dests {
		destinationDtos = append(destinationDtos, *mapper.MapDestinationToDto(d))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, destinationDtos)
}

func (h *DestinationHandler) GetDestinationByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	dest, err := h.service.GetDestinationByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "destination not found"})
	}
	return c.JSON(http.StatusOK, dest)
}

func (h *DestinationHandler) CreateDestination(c echo.Context) error {
	var dto dto.CreateDestinationRequestDto
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	created, err := h.service.CreateDestination(mapper.MapFromCreateReq(dto))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *DestinationHandler) DeleteDestination(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.service.DeleteDestination(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
