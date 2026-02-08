package route

import (
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Group) {
	repository := repository.NewUserRepository()
	userService := service.NewUserService(repository)
	authService := service.NewAuthService(repository)

	handler := handler.NewUserHandler(userService, authService)

	e.GET("/users", handler.GetAllUsers)
	e.GET("/users/:id", handler.GetUserByID)
	e.POST("/users", handler.CreateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
	e.POST("/users/login", handler.Login)
	e.POST("/users/register", handler.Register)
}
