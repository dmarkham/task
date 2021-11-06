package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func startServer() {
	http.HandleFunc("/", mainHandler)
	log.Panicln(http.ListenAndServe(":5000", nil))
}

func main() {
	startServer()
}
