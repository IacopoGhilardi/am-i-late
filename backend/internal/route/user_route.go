package route

import (
	"github.com/iacopoGhilardi/amILate/internal/app"
	"github.com/iacopoGhilardi/amILate/internal/middleware"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Group, a *app.App) {
	handler := a.UserHandler

	e.POST("/users/login", handler.Login)
	e.POST("/users/register", handler.Register)
	e.POST("/users/forgot-password", handler.ForgotPassword)
	e.POST("/users/reset-password", handler.ResetPassword)

	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/users", handler.GetAllUsers)
	protected.GET("/users/:id", handler.GetUserByID)
	protected.POST("/users", handler.CreateUser)
	protected.DELETE("/users/:id", handler.DeleteUser)
}
