package config

import (
	"os"
	"testing"
)

func TestReadFromEnv(t *testing.T) {
	os.Setenv("INFLUXDB_URL", "http://localhost:8086")
	os.Setenv("INFLUXDB_USERNAME", "admin")
	os.Setenv("INFLUXDB_PASSWORD", "admin")

	ReadFromEnv()
}

func TestGetConfig(t *testing.T) {
	config := GetConfig()

	if config.InfluxdbUrl != "http://localhost:8086" {
		t.Errorf("InfluxdbURL is not set correctly")
	}

	if config.InfluxdbUsername != "admin" {
		t.Errorf("InfluxUsername is not set correctly")
	}

	if config.InfluxdbPassword != "admin" {
		t.Errorf("InfluxPassword is not set correctly")
	}

}
