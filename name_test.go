package main

import (
	"reflect"
	"testing"
)

func TestGetJokeUrl(t *testing.T) {
	testNames := []Name{
		{"Edward", "Wood"},
		{"Dalila", "Fischetti"},
		{"Candida", "Prociuk"},
		{"Chiavana", "Weidert"},
		{"Kay", "Roura"},
		{"Firas", "Quibodeaux"},
		{"Abdelghafour", "Schoewe"},
		{"Anilda", "Fairburn"},
		{"Jatinder", "Dayao"},
		{"Jiawei", "Tretina"},
		{"Hagar", "Buchannon"},
		{"Arsalan", "Jandrin"},
	}
	for _, name := range testNames {
		url := name.getJokeUrl()
		result := getNameFromUrl([]rune(url))
		if !reflect.DeepEqual(name, result) {
			t.Fatalf("Expected: %v, got: %v", name, result)
		}
	}

}

func getNameFromUrl(url []rune) Name {
	var firstname, lastname string
	firstnameDone := false
	startLastName := false
	for i := 44; i < len(url); i++ {
		if url[i] == '&' {
			if startLastName {
				break
			}
			firstnameDone = true
		} else if url[i] == '=' && firstnameDone {
			startLastName = true
			continue
		}

		if !firstnameDone {
			firstname += string(url[i])
		} else if startLastName {
			lastname += string(url[i])
		}
	}
	return Name{firstname, lastname}
}
