package main

import (
	"go_server/microservice"
	"log"
	"os"
)

func main() {
	server, err := BuildServer()
	if err != nil {
		log.Println("Error")
		os.Exit(1)
	}

	err = microservice.Start(server)
	if err != nil {
		log.Println("Fatal error starting service: ", "error", err)
		os.Exit(1)
	}
}
