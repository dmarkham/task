package main

import "encoding/json"

const nameUrl = "https://names.mcquay.me/api/v0/"

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// getJokeUrl() makes the url for the joke request
// using the given Name.
func (name *Name) getJokeUrl() string {
	return "http://api.icndb.com/jokes/random?firstName=" + name.FirstName + "&lastName=" + name.LastName + "&limitTo=nerdy"
}

// TODO: add input validation for cases where it doesn't match the expected json.
// BytesToName() takes bytes from a request and returns
// a Name and error.
func bytesToName(rawName []byte) (Name, error) {
	name := Name{}
	err := json.Unmarshal(rawName, &name)
	return name, err
}
