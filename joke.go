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

type jokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

func fetchJoke() Joke {

	client := http.Client{
		Timeout: TIMEOUT,
	}

	res, reqErr := client.Get(JOKE_URL)
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

	jokeResponse := jokeResponse{}

	err := json.Unmarshal(body, &jokeResponse)

	if err != nil {
		fmt.Println("Joke unmarshal error:", err)
		return Joke{}
	}

	return jokeResponse.Value
}
