package main

import (
	"fmt"
	"log"

	api "github.com/bus710/task/src/api"
)

func main() {
	fmt.Println("Hello")

	fmt.Println("Bye")

	apiServer, err := api.SetupAPIServer()
	if err != nil {
		log.Println("API server setup was not successful, abort")
	}

	apiServer.Run(":5000")
}
