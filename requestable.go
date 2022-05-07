package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Requestable interface {
	GetUrl() string
}

func makeRequest(r Requestable) ([]byte, error) {
	res, err := http.Get(r.GetUrl())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func toStruct(data []byte, buf any) error {
	return json.Unmarshal(data, buf)
}

func Process(r Requestable) error {
	res, err := makeRequest(r)
	if err != nil {
		return err
	}
	return toStruct(res, r)
}
