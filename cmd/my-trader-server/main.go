package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	grpcAdapter "github.com/mukappalambda/my-trader/internal/adapters/grpc"
	restAdapter "github.com/mukappalambda/my-trader/internal/adapters/rest"
)

var (
	port    = flag.Int("port", 50051, "server port")
	ginMode = flag.String("gin-mode", "release", "gin mode (debug|release)")
)

func main() {
	flag.Parse()
	dsn := os.Getenv("DATABASE_URL")
	srv, err := grpcAdapter.NewGrpcServer(dsn)
	if err != nil {
		log.Fatal(err)
	}
	schemaRegistryServer, err := restAdapter.NewSchemaRegistryServer(*ginMode)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	registryPort := 8081
	go func() {
		log.Printf("schema registry server is listening at port: %d\n", registryPort)
		if err := schemaRegistryServer.Run(fmt.Sprintf("localhost:%d", registryPort)); err != nil {
			log.Fatal(err)
		}
	}()
	if err := srv.Run(*port); err != nil {
		log.Fatal(err)
	}
}
