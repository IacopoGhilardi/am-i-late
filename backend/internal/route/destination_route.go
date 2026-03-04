package route

import (
	"github.com/iacopoGhilardi/amILate/internal/app"
	"github.com/iacopoGhilardi/amILate/internal/middleware"
	"github.com/labstack/echo/v4"
)

func InitDestinationRoutes(e *echo.Group, a *app.App) {
	handler := a.DestinationHandler

	e.Use(middleware.JWTMiddleware())
	e.GET("/destinations", handler.GetAllDestinations)
	e.GET("/destinations/:id", handler.GetDestinationByID)
	e.POST("/destinations", handler.CreateDestination)
	e.DELETE("/destinations/:id", handler.DeleteDestination)
	e.GET("destinations/geocode", handler.GeocodeDestination)
}
