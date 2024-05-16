package main

import (
	"flag"
	"log"

	"github.com/mukappalambda/my-trader/pkg/server"
)

var (
	addr = flag.String("addr", ":8080", "addr")
)

func main() {
	flag.Parse()

	if err := server.Run(*addr); err != nil {
		log.Fatal(err)
	}
}
