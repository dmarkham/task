package requests

const nameUrl = "https://names.mcquay.me/api/v0/"

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (n *Name) GetUrl() string {
	return nameUrl
}
