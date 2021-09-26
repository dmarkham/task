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

	res = textProcessing(res)

	c.String(http.StatusOK, res)
}

func textProcessing(res string) string {
	// Need to remove a space right before/after of the name
	resArr := strings.Split(res, " ")
	res = ""
	for _, s := range resArr {
		fmt.Println(s)
		s = strings.Replace(s, " ", "", -1)
		fmt.Println("=>", s)

		// Need to avoid double spaces
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
	res = res[1 : len(res)-1]

	// Need to add a new line at the end of the string to remove \r
	res = fmt.Sprint(res, "\n")

	return res
}
