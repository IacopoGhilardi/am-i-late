package destination

import (
	"github.com/labstack/echo/v4"
)

func InitDestinationRoutes(e *echo.Group) {
	service := NewDestinationService()
	handler := NewDestinationHandler(service)

	e.GET("/destinations", handler.GetAllDestinations)
	e.GET("/destinations/:id", handler.GetDestinationByID)
	e.POST("/destinations", handler.CreateDestination)
	e.DELETE("/destinations/:id", handler.DeleteDestination)
}
