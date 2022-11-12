package main

import "time"

const (
	TIMEOUT  = time.Millisecond * 200
	NAME_URL = "http://localhost:9005/name"
	JOKE_URL = "http://localhost:9005/joke"
	PORT     = ":8080"
)
