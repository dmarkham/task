package main

const nameUrl = "https://names.mcquay.me/api/v0/"

type name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (n *name) GetUrl() string {
	return nameUrl
}
