package db

import (
	log "github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(connectionString string) {
	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Error while connecting to db: %v", err)
	}
	log.Info("Postgres DB connected")
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
		log.Fatal("❌ Db not initialized, call db.Connect()")
	}
	return DB
}
