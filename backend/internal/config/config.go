package config

import (
	"time"

	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT" default:"1323"`
	DbUrl      string `mapstructure:"DB_URL"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`

	//Google maps
	GoogleMapsApiKey  string        `mapstructure:"GOOGLE_MAPS_API_KEY"`
	GoogleMapsBaseUrl string        `mapstructure:"GOOGLE_MAPS_API_BASE_URL"`
	GoogleMapsTimeout time.Duration `mapstructure:"GOOGLE_MAPS_API_TIMEOUT" default:"10000"`

	// Email
	ResendApiKey     string `mapstructure:"RESEND_API_KEY"`
	EmailFromAddress string `mapstructure:"EMAIL_FROM_ADDRESS"`
	ResetPasswordURL string `mapstructure:"RESET_PASSWORD_URL"`
}

func LoadConfig(path string) (*Config, error) {
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
		return &config, err
	}

	return &config, nil
}
