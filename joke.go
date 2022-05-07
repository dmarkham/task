package main

import (
	"sync"
	"task/requests"
)

type Joke struct {
	*requests.Name
	*requests.Content
}

func (j *Joke) processNameAndContent() error {
	err := requests.Do(j.Name)
	if err != nil {
		return err
	}
	err = requests.Do(j.Content)
	if err != nil {
		return err
	}
	return nil
}

func (j *Joke) DoOne(wg *sync.WaitGroup, r requests.Requestable) error {
	defer wg.Done()
	err := requests.Do(r)
	if err != nil {
		return err
	}
	return nil
}

func NewJoke() string {
	j := Joke{
		Name:    &requests.Name{},
		Content: &requests.Content{},
	}
	err := j.processNameAndContent()
	if err != nil {
		return err.Error()
	}

	return j.SwapPlaceholders(j.Name)
}
