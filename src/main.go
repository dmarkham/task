// Fetching a name
// $ curl "https://names.mcquay.me/api/v0/"
// {“first_name”:“Hasina”,“last_name”:“Tanweer”}
//
// Fetching a joke
// $ curl "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"
// { “type”: “success”, “value”: { “id”: 181, “joke”: “John Doe’s OSI network model has only one layer - Physical.“, “categories”: [“nerdy”] } }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")

	fmt.Println("Bye")

	apiServer, err := setupAPIServer()
	if err != nil {
		log.Println("API server setup was not successful, abort")
	}

	apiServer.Run(":5000")
}

func setupAPIServer() (*gin.Engine, error) {
	apiServer := gin.New()
	apiServer.Use(gin.Logger())

	apiServer.GET("/", taskHandler)

	return apiServer, nil
}

func taskHandler(c *gin.Context) {
	log.Println("task hit")
	c.String(http.StatusOK, "task hit\n")
}
