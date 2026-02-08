package config

import (
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT" default:"1323"`
	DbUrl      string `mapstructure:"DB_URL"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`

	//Tom Tom
	TomTomApiKey  string `mapstructure:"REST_CLIENT_TOMTOM_API_KEY"`
	TomTomBaseUrl string `mapstructure:"REST_CLIENT_TOMTOM_BASE_URL"`
	TomTomTimeout int    `mapstructure:"REST_CLIENT_TOMTOM_TIMEOUT" default:"10"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("No config file found, using environment variables: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		logger.Error("Failed to unmarshal config: %v", err)
		return config, err
	}

	return config, nil
}
