package main

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/config"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/route"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config: %v", err)
	}
	logger.Info("config loaded: %+v", cfg)

	security.SetJWTSecret(cfg.JwtSecret)
	db.Connect(cfg.DbUrl)
	echoInstance := echo.New()

	apiV1 := echoInstance.Group("/api/v1")

	//Commons
	commons.InitUtilityRoute(apiV1)

	//User
	route.InitUserRoutes(apiV1)

	//Destination
	route.InitDestinationRoutes(apiV1)
	route.InitAppointmentRoutes(apiV1)

	echoInstance.Logger.Fatal(echoInstance.Start(":" + cfg.ServerPort))
}
