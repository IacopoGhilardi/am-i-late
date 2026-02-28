package route

import (
	"github.com/iacopoGhilardi/amILate/internal/app"
	"github.com/iacopoGhilardi/amILate/internal/middleware"
	"github.com/labstack/echo/v4"
)

func InitAppointmentRoutes(e *echo.Group, a *app.App) {
	handler := a.AppointmentHandler

	e.Use(middleware.JWTMiddleware())
	e.GET("/appointments", handler.GetAllAppointments)
	e.GET("/appointments/:id", handler.GetAppointmentByPublicId)
	e.POST("/appointments", handler.CreateAppointment)
	e.DELETE("/appointments/:id", handler.DeleteAppointment)
}
