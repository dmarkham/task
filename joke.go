package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	Categories []string `json:"categories"`
	Id         int      `json:"id"`
	Joke       string   `json:"joke"`
}

// Response from the joke server has an additional type property. I have only received the success type. TODO: add additional checking for other types.
type jokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

func fetchJoke(name Name) Joke {
	jokeUrl := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", name.FirstName, name.LastName)

	client := http.Client{
		Timeout: TIMEOUT,
	}

	res, reqErr := client.Get(jokeUrl)
	if reqErr != nil {
		log.Fatal("Joke request error:", reqErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal("Joke read error:", readErr)
	}

	// The URL retrieves a joke response and the value property of that response is the joke object
	jokeResponse := jokeResponse{}

	err := json.Unmarshal(body, &jokeResponse)

	if err != nil {
		fmt.Println("Joke unmarshal error:", err)
		return Joke{}
	}

	// Extracting the joke object from the response
	return jokeResponse.Value
}
