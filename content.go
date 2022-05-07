package main

const firstnamePlaceholder = "FIRST_NAME"
const lastnamePlaceholder = "LAST_NAME"
const jokeUrl = "http://api.icndb.com/jokes/random?firstName=" + firstnamePlaceholder + "&lastName=" + lastnamePlaceholder + "&limitTo=nerdy"

type content struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func (c *content) GetUrl() string {
	return jokeUrl
}

func NewContent() *content {
	return &content{}
}
