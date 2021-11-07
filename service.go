package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// mainHandler() responds to requests to /* with the joke.
func mainHandler(w http.ResponseWriter, _ *http.Request) {
	name, err := getName()
	if err != nil {
		log.Panicln(err)
	}
	joke, err := name.getJoke()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Fprintf(w, joke.Value.Joke)
}

// getName() calls the name api, returning a Name and error.
func getName() (Name, error) {
	res, err := http.Get(nameUrl)
	if err != nil {
		return Name{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Name{}, err
	}
	return bytesToName(body)
}

// getJoke() gets the joke for a given Name.
func (name *Name) getJoke() (Joke, error) {
	jokeUrl := name.getJokeUrl()
	res, err := http.Get(jokeUrl)
	if err != nil {
		return Joke{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Joke{}, err
	}
	return bytesToJoke(body)
}

// TODOs for full production ready:
// Depending on the intended response for a rate limit, possibly
// respond with a result from an archived joke + name

// startServer() specifies the handle functions and port and starts the server.
func startServer() {
	http.HandleFunc("/", mainHandler)
	log.Panicln(http.ListenAndServe(":5000", nil))
}

func main() {
	startServer()
}
