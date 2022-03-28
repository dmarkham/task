package api

type NamesBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type JokesBody struct {
	Type       string `json:"type"`
	Value      JokesBodyValue
	Vategories []string
}

type JokesBodyValue struct {
	ID   uint   `json:"id"`
	Joke string `json:"joke"`
}
