package main

import (
	"flag"
	"log"
	"os"

	grpcAdapter "github.com/mukappalambda/my-trader/internal/adapters/grpc"
)

var port = flag.Int("port", 50051, "server port")

func main() {
	flag.Parse()
	dsn := os.Getenv("DATABASE_URL")
	srv, err := grpcAdapter.NewGrpcServer(dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(*port); err != nil {
		log.Fatal(err)
	}
}
