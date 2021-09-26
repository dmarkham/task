package main

import (
	"fmt"
	"log"

	api "github.com/bus710/task/src/api"
)

func main() {
	fmt.Println("Hello")

	apiServer, err := api.SetupAPIServer()
	if err != nil {
		log.Println("API server setup was not successful, abort")
	}

	apiServer.Run(":5000")

	fmt.Println("Bye")
}
