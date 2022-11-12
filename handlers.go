package main

import (
	"fmt"
	"net/http"
	"strings"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	name := fetchName()
	joke := fetchJoke()

	// Replacing John Doe from the joke with the given name from the response
	response := strings.Replace(joke.Joke, "John Doe", fmt.Sprintf("%s %s", name.FirstName, name.LastName), -1)

	// response must be in an array of bytes
	b := []byte(response)

	w.Write(b)
}
