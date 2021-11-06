package main

import "encoding/json"

const nameUrl = "https://names.mcquay.me/api/v0/"

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (name *Name) getJokeUrl() string {
	return "http://api.icndb.com/jokes/random?firstName=" + name.FirstName + "&lastname=" + name.LastName + "&limitTo=nerdy"
}

func bytesToName(rawName []byte) (Name, error) {
	name := Name{}
	err := json.Unmarshal(rawName, &name)
	return name, err
}
