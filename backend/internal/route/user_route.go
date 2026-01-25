package route

import (
	handler2 "github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Group, service *service.UserService) {
	handler := handler2.NewUserHandler(service)

	e.GET("/users", handler.GetAllUsers)
	e.GET("/users/:id", handler.GetUserByID)
	e.POST("/users", handler.CreateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
	e.POST("/users/login", handler.Login)
	e.POST("/users/register", handler.Register)
}
