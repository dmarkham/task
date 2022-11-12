package main

import (
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch name and defer until the name returns
	name := fetchName()
	// Use name to replace John Doe when fetching the joke
	joke := fetchJoke(name)

	// response must be in an array of bytes
	b := []byte(joke.Joke)

	w.Write(b)
}
