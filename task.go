package main

import (
	"log"
	"net/http"
)

func main() {
	// Handling only the default route and a single request at a time.
	// TODO: add more routes as needed. Add additional logic for cuncurent requests and load balancers.
	http.HandleFunc("/", TaskHandler)
	log.Println("listen on port: ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
