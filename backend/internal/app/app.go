package app

import (
	"github.com/iacopoGhilardi/amILate/internal/client/googleMaps"
	"github.com/iacopoGhilardi/amILate/internal/config"
	"github.com/iacopoGhilardi/amILate/internal/handler"
	"github.com/iacopoGhilardi/amILate/internal/repository"
	"github.com/iacopoGhilardi/amILate/internal/service"
	log "github.com/iacopoGhilardi/amILate/internal/utils/logger"
)

type App struct {
	// Handlers
	UserHandler        *handler.UserHandler
	AppointmentHandler *handler.AppointmentHandler
	DestinationHandler *handler.DestinationHandler

	// Services
	AppointmentService *service.AppointmentService
	UserService        *service.UserService
	DestinationService *service.DestinationService
	AuthService        *service.AuthService
	MapService         *service.MapService
}

func NewApp(cfg *config.Config) *App {
	userRepo := repository.NewUserRepository()
	appointmentRepo := repository.NewAppointmentRepository()
	destinationRepo := repository.NewDestinationRepository()

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)
	destinationService := service.NewDestinationService(destinationRepo)

	mapsClient, err := googleMaps.NewClient(cfg.GoogleMapsApiKey)
	if err != nil {
		log.Fatal("Error while initializing maps client", err)
	}

	mapService := service.NewMapService(mapsClient)

	return &App{
		UserHandler:        handler.NewUserHandler(userService, authService),
		AppointmentHandler: handler.NewAppointmentHandler(appointmentService),
		DestinationHandler: handler.NewDestinationHandler(destinationService),
		AppointmentService: appointmentService,
		UserService:        userService,
		DestinationService: destinationService,
		AuthService:        authService,
		MapService:         mapService,
	}
}
