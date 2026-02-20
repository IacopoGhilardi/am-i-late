package route

import (
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/middleware"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/service"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Group) {
	repository := repository.NewUserRepository()
	userService := service.NewUserService(repository)
	authService := service.NewAuthService(repository)

	handler := handler.NewUserHandler(userService, authService)

	e.POST("/users/login", handler.Login)
	e.POST("/users/register", handler.Register)

	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/users", handler.GetAllUsers)
	protected.GET("/users/:id", handler.GetUserByID)
	protected.POST("/users", handler.CreateUser)
	protected.DELETE("/users/:id", handler.DeleteUser)
}
