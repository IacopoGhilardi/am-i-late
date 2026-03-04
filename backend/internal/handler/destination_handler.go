package handler

import (
	"net/http"
	"strconv"

	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	_interface "github.com/iacopoGhilardi/amILate/internal/service/interface"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/labstack/echo/v4"
)

type DestinationHandler struct {
	service     _interface.DestinationServiceInterface
	mapsService _interface.MapsServiceInterface
}

func NewDestinationHandler(service _interface.DestinationServiceInterface, mapService _interface.MapsServiceInterface) *DestinationHandler {
	return &DestinationHandler{service: service, mapsService: mapService}
}

func (h *DestinationHandler) GetAllDestinations(c echo.Context) error {
	logger.Info("Getting all destinations")
	dests, err := h.service.GetAllDestinations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	var destinationDtos []dto.DestinationDto
	for _, d := range dests {
		destinationDtos = append(destinationDtos, *mapper.MapDestinationToDto(d))
	}
	return c.JSON(http.StatusOK, commons.Success(destinationDtos))
}

func (h *DestinationHandler) GetDestinationByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail("invalid id"))
	}
	dest, err := h.service.GetDestinationByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, commons.Fail("destination not found"))
	}
	return c.JSON(http.StatusOK, commons.Success(dest))
}

func (h *DestinationHandler) CreateDestination(c echo.Context) error {
	var createDto dto.CreateDestinationRequestDto
	if err := c.Bind(&createDto); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail(err.Error()))
	}
	created, err := h.service.CreateDestination(mapper.MapFromCreateReq(createDto))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.JSON(http.StatusCreated, commons.Success(created))
}

func (h *DestinationHandler) DeleteDestination(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.Fail("invalid id"))
	}
	if err := h.service.DeleteDestination(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *DestinationHandler) GeocodeDestination(c echo.Context) error {
	address := c.QueryParam("address")
	if address == "" {
		return c.JSON(http.StatusBadRequest, commons.Fail("address is required"))
	}
	results, err := h.mapsService.Geocode(c.Request().Context(), address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Fail(err.Error()))
	}
	return c.JSON(http.StatusOK, commons.Success(results))
}
