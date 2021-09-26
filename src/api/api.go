package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupAPIServer() (*gin.Engine, error) {
	apiServer := gin.New()
	apiServer.Use(gin.Logger())

	apiServer.GET("/", taskHandler)

	return apiServer, nil
}

// taskHandler gets data from 2 differnt sources and pushes a mixed string between the two to clients.
func taskHandler(c *gin.Context) {
	// 1. Fetching a name from "https://names.mcquay.me/api/v0/"
	// => {“first_name”:“Hasina”,“last_name”:“Tanweer”}

	namesAPIUrl := "https://names.mcquay.me/api/v0/"
	namesBody := NamesBody{}

	resp, err := http.Get(namesAPIUrl)
	if resp == nil || err != nil {
		log.Println("Error:", err)
		c.String(http.StatusNoContent, "Call to names API failed")
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&namesBody)

	// 2. Fetching a joke from "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"
	// => “type”: “success”, “value”: { “id”: 181, “joke”: “John Doe’s OSI network model has only one layer - Physical.“, “categories”: [“nerdy”] } }

	jokesAPIUrl := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=&limitTo=nerdy", namesBody.FirstName)
	jokesBody := JokesBody{}

	resp, err = http.Get(jokesAPIUrl)
	if resp == nil || err != nil {
		log.Println("Error:", err)
		c.String(http.StatusNoContent, "Call to jokes API failed")
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&jokesBody)
	res := jokesBody.Value.Joke
	res = strings.Replace(res, " ", "", 1) // Need to remove a space right after the name
	res = fmt.Sprint(res, "\n")
	// TODO: possibly need to match the pronounces and names' gender?

	c.String(http.StatusOK, res)
}
