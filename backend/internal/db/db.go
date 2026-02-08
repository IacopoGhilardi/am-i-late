package db

import (
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(connectionString string) {
	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Fatal("Error while connecting to db: %v", err)
	}
	logger.Info("Postgres DB connected")
}

func Ping() error {
	sqlDB, err := GetDB().DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func GetDB() *gorm.DB {
	if DB == nil {
		logger.Fatal("❌ Db not initialized, call db.Connect()")
	}
	return DB
}
