package influxdb

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type ClientConfig struct {
	Url      string
	Username string
	Password string
}

func CreateClient(cc ClientConfig) influxdb2.Client {

	autoToken := fmt.Sprintf("%s:%s", cc.Username, cc.Password)
	client := influxdb2.NewClient(cc.Url, autoToken)

	return client
}
