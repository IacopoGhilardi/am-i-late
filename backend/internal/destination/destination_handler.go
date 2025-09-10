package destination

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DestinationHandler struct {
	service *DestinationService
}

func NewDestinationHandler(service *DestinationService) *DestinationHandler {
	return &DestinationHandler{service: service}
}

func (h *DestinationHandler) GetAllDestinations(c echo.Context) error {
	dests, err := h.service.GetAllDestinations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, dests)
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
	var d Destination
	if err := c.Bind(&d); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	created, err := h.service.CreateDestination(&d)
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
