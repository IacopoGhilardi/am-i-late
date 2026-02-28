package app

import (
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/service"
)

type App struct {
	// Handlers
	UserHandler        *handler.UserHandler
	AppointmentHandler *handler.AppointmentHandler
	DestinationHandler *handler.DestinationHandler

	//Services
	AppointmentService *service.AppointmentService
	UserService        *service.UserService
	DestinationService *service.DestinationService
	AuthService        *service.AuthService
}

func NewApp() *App {
	userRepo := repository.NewUserRepository()
	appointmentRepo := repository.NewAppointmentRepository()
	destinationRepo := repository.NewDestinationRepository()

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)
	destinationService := service.NewDestinationService(destinationRepo)

	return &App{
		UserHandler:        handler.NewUserHandler(userService, authService),
		AppointmentHandler: handler.NewAppointmentHandler(appointmentService),
		DestinationHandler: handler.NewDestinationHandler(destinationService),
		AppointmentService: appointmentService,
		UserService:        userService,
		DestinationService: destinationService,
		AuthService:        authService,
	}
}
