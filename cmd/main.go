package main

import (
	"exec-processor/internal/server"
	"log"
)

func main() {
	var pathToSettings = "server_config.yaml"
	server := server.New(pathToSettings)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
