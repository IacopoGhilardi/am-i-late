package alarm

import (
	"github.com/labstack/echo/v4"
)

func InitAlarmRoutes(e *echo.Group) {
	service := NewAlarmService()
	handler := NewAlarmHandler(service)

	e.GET("/alarms", handler.GetAllAlarms)
	e.GET("/alarms/:id", handler.GetAlarmByID)
	e.POST("/alarms", handler.CreateAlarm)
	e.DELETE("/alarms/:id", handler.DeleteAlarm)
}
