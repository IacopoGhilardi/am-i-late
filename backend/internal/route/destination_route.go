package route

import (
	handler2 "github.com/iacopoGhilardi/amILate/internal/handler"
	service2 "github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/labstack/echo/v4"
)

func InitDestinationRoutes(e *echo.Group) {
	service := service2.NewDestinationService()
	handler := handler2.NewDestinationHandler(service)

	e.GET("/destinations", handler.GetAllDestinations)
	e.GET("/destinations/:id", handler.GetDestinationByID)
	e.POST("/destinations", handler.CreateDestination)
	e.DELETE("/destinations/:id", handler.DeleteDestination)
}
