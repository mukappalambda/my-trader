package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	InfluxdbUrl      string
	InfluxdbUsername string
	InfluxdbPassword string
}

// define a global config struct
var config *Config

func GetConfig() *Config {
	if config == nil {
		panic("config not initialized")
	}

	return config
}

func ReadFromEnv() error {
	if config != nil {
		return errors.New("config already initialized")
	}

	err := godotenv.Load()

	if err != nil {
		return err
	}

	newConfig := new(Config)
	newConfig.InfluxdbUrl = os.Getenv("INFLUXDB_URL")
	newConfig.InfluxdbUsername = os.Getenv("INFLUXDB_USERNAME")
	newConfig.InfluxdbPassword = os.Getenv("INFLUXDB_PASSWORD")

	config = newConfig

	return nil
}
