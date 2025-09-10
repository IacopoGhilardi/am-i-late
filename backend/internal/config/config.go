package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT" default:"1323"`
	DbUrl      string `mapstructure:"DB_URL"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, using environment variables: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Failed to unmarshal config: %v", err)
		return config, err
	}

	return config, nil
}
