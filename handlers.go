package main

import (
	"fmt"
	"net/http"
	"strings"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	name := fetchName()
	joke := fetchJoke()

	response := strings.Replace(joke.Joke, "John Doe", fmt.Sprintf("%s %s", name.FirstName, name.LastName), -1)
	b := []byte(response)

	w.Write(b)
}
