package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func fetchName() Name {

	client := http.Client{
		Timeout: TIMEOUT,
	}

	res, reqErr := client.Get(NAME_URL)
	if reqErr != nil {
		log.Fatal("Name request error:", reqErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal("Name read error:", readErr)
	}

	name := Name{}

	err := json.Unmarshal(body, &name)

	if err != nil {
		fmt.Println("Name unmarshal error:", err)
		return Name{}
	}

	return name
}
