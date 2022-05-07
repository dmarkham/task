package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getJoke(c *gin.Context) {
	got := NewJoke()
	c.IndentedJSON(http.StatusOK, got)
}

func main() {
	router := gin.Default()
	router.GET("/", getJoke)
	router.Run(":8080")
}
