package requests

import "strings"

const firstnamePlaceholder = "FIRST_NAME"
const lastnamePlaceholder = "LAST_NAME"
const jokeUrl = "http://api.icndb.com/jokes/random?firstName=" + firstnamePlaceholder + "&lastName=" + lastnamePlaceholder + "&limitTo=nerdy"

type Content struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func (c *Content) GetUrl() string {
	return jokeUrl
}

func (c *Content) SwapPlaceholders(n *Name) string {
	temp := strings.ReplaceAll(c.Value.Joke, firstnamePlaceholder, n.FirstName)
	return strings.ReplaceAll(temp, lastnamePlaceholder, n.LastName)
}
