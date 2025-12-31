package user

import (
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Group, service *UserService) {
	handler := NewUserHandler(service)

	e.GET("/users", handler.GetAllUsers)
	e.GET("/users/:id", handler.GetUserByID)
	e.POST("/users", handler.CreateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
	e.POST("/users/login", handler.Login)
}
