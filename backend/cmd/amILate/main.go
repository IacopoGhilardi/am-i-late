package main

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/config"
	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/iacopoGhilardi/amILate/internal/route"
	"github.com/iacopoGhilardi/amILate/internal/service"
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
	service := service.NewUserService()
	route.InitUserRoutes(apiV1, service)

	//Destination
	route.InitDestinationRoutes(apiV1)

	//Todo: Rotte

	echoInstance.Logger.Fatal(echoInstance.Start(":" + cfg.ServerPort))
}
