package main

import (
	"fmt"
	"strings"
	"sync"
)

type Joke struct {
	*name
	*content
	result string
}

func (j *Joke) swapPlaceholders() {
	j.result = strings.ReplaceAll(j.content.Value.Joke, firstnamePlaceholder, j.FirstName)
	j.result = strings.ReplaceAll(j.result, lastnamePlaceholder, j.LastName)
}

func (j *Joke) processNameAndContent() error {
	err := Process(j.name)
	if err != nil {
		return err
	}
	err = Process(j.content)
	if err != nil {
		return err
	}
	return nil
}

func (j *Joke) DoOne(wg *sync.WaitGroup, r Requestable) error {
	defer wg.Done()
	err := Process(r)
	if err != nil {
		return err
	}
	return nil
}

func NewJoke() string {
	j := Joke{
		name:    &name{},
		content: &content{},
		result:  "",
	}
	err := j.processNameAndContent()
	if err != nil {
		return err.Error()
	}

	fmt.Printf("Joke before: %v\n", j.content.Value.Joke)
	fmt.Printf("Joke url: %v\n", j.content.GetUrl())
	fmt.Printf("Name: %v\n", j.FirstName+" "+j.LastName)
	fmt.Printf("Joke after: %v\n", j.result)

	j.swapPlaceholders()
	return j.result
}
