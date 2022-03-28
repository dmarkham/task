package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupAPIServer returns Gonic HTTP engin instance with handlers configured
func SetupAPIServer() (*gin.Engine, error) {
	apiServer := gin.New()
	apiServer.Use(gin.Logger())

	apiServer.GET("/", taskHandler)

	return apiServer, nil
}

// taskHandler gets data from 2 differnt sources and pushes a mixed string between the two to clients.
// TODO: to be tenacious and stable, we may add short-circuit and retry patterns.
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

	res = textProcessing(res)

	c.String(http.StatusOK, res)
}

// textProcessing trims and replces to correct possible errors from the input string
func textProcessing(res string) string {
	// Need to remove a space right before/after of the name
	resArr := strings.Split(res, " ")
	res = ""
	for _, s := range resArr {
		// Need to remove white space from head/tail
		s = strings.Replace(s, " ", "", -1)

		// Sometimes there is a whitespace. Need to remove
		if len(s) < 2 {
			continue
		}

		// Need to remove the second s from possessive ('ss => 's)
		if s == "'ss" {
			s = "'s"
		}

		// Need to remove a spcae between the name and possessive ('s)
		if s[0] != '\'' {
			res = fmt.Sprint(res, " ", s)
		} else {
			res = fmt.Sprint(res, "", s)
		}
	}

	// Need to remove if the first character is whitespace
	if res[0] == ' ' {
		res = res[1 : len(res)-1]
	}

	// Need to add a new line at the end of the string to remove \r
	res = fmt.Sprint(res, "\n")

	return res
}
