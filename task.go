package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", TaskHandler)
	log.Println("listen on port: ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
