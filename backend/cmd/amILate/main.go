package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/iacopoGhilardi/amILate/internal/app"
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/config"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/route"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	customValidator "github.com/iacopoGhilardi/amILate/internal/utils/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config: %v", err)
	}
	logger.Info("config loaded: %+v", cfg)

	a := app.NewApp(cfg)

	security.SetJWTSecret(cfg.JwtSecret)
	db.Connect(cfg.DbUrl)
	echoInstance := echo.New()
	echoInstance.Validator = &customValidator.CustomValidator{Validator: validator.New()}

	apiV1 := echoInstance.Group("/api/v1")

	// Routes
	commons.InitUtilityRoute(apiV1)
	route.InitUserRoutes(apiV1, a)
	route.InitDestinationRoutes(apiV1, a)
	route.InitAppointmentRoutes(apiV1, a)

	echoInstance.Logger.Fatal(echoInstance.Start(":" + cfg.ServerPort))
}
