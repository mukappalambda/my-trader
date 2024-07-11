package main

import (
	"flag"
	"log"

	grpcAdapter "github.com/mukappalambda/my-trader/internal/adapters/grpc"
)

var port = flag.Int("port", 50051, "server port")

func main() {
	flag.Parse()
	app := grpcAdapter.NewApp()
	if err := app.Run(*port); err != nil {
		log.Fatal(err)
	}
}
