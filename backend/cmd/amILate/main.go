package main

import (
	"github.com/iacopoGhilardi/amILate/internal/alarm"
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/config"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/destination"
	"github.com/iacopoGhilardi/amILate/internal/user"
	"github.com/iacopoGhilardi/amILate/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config: %v", err)
	}

	logger.Info("config loaded: %+v", cfg)

	db.Connect(cfg.DbUrl)
	echoInstance := echo.New()

	apiV1 := echoInstance.Group("/api/v1")

	//Commons
	commons.InitUtilityRoute(apiV1)

	//User
	service := user.NewUserService()
	user.InitUserRoutes(apiV1, service)

	//Alarm
	alarm.InitAlarmRoutes(apiV1)

	//Destination
	destination.InitDestinationRoutes(apiV1)

	echoInstance.Logger.Fatal(echoInstance.Start(":" + cfg.ServerPort))
}
