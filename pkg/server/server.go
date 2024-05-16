package server

import (
	"fmt"

	"github.com/mukappalambda/my-trader/pkg/config"
	"github.com/mukappalambda/my-trader/pkg/web"
	"github.com/sirupsen/logrus"
)

func Run(addr string) error {
	// read config from env
	logrus.Info("Starting server...")

	err := config.ReadFromEnv()
	if err != nil {
		return fmt.Errorf("error reading config: %q", err)
	}

	c := config.GetConfig()

	logrus.Infof("Load InfluxdbUrl: %s", c.InfluxdbUrl)
	logrus.Infof("Load InfluxdbUsername: %s", c.InfluxdbUsername)
	logrus.Infof("Load InfluxdbPassword: %s", c.InfluxdbPassword)

	logrus.Info("Reading config from env...")

	app := web.NewApiServer()
	return app.Listen(addr)
}
