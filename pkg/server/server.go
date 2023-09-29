package server

import (
	"github.com/mukappalambda/my-trader/pkg/config"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	// read config from env
	logrus.Info("Starting server...")

	err := config.ReadFromEnv()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	c := config.GetConfig()

	logrus.Infof("Load InfluxdbUrl: %s", c.InfluxdbUrl)
	logrus.Infof("Load InfluxdbUsername: %s", c.InfluxdbUsername)
	logrus.Infof("Load InfluxdbPassword: %s", c.InfluxdbPassword)

	logrus.Info("Reading config from env...")

}