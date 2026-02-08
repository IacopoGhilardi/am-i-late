package route

import (
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/labstack/echo/v4"
)

func InitAppointmentRoutes(e *echo.Group) {
	repository := repository.NewAppointmentRepository()
	service := service.NewAppointmentService(repository)
	handler := handler.NewAppointmentHandler(service)

	e.GET("/appointments", handler.GetAllAppointments)
	e.GET("/appointments/:id", handler.GetAppointmentByPublicId)
	e.POST("/appointments", handler.CreateAppointment)
	e.DELETE("/appointments/:id", handler.DeleteAppointment)
}
