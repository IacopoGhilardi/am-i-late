package route

import (
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/labstack/echo/v4"
)

func InitDestinationRoutes(e *echo.Group) {
	repository := repository.NewDestinationRepository()
	service := service.NewDestinationService(repository)
	handler := handler.NewDestinationHandler(service)

	e.GET("/destinations", handler.GetAllDestinations)
	e.GET("/destinations/:id", handler.GetDestinationByID)
	e.POST("/destinations", handler.CreateDestination)
	e.DELETE("/destinations/:id", handler.DeleteDestination)
}
