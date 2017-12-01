package main

import (
	"encoding/json"
	"io/ioutil"
)

type AppParams struct {
	NodeAddress   string
	QuoteDuration int
}

func readParams() AppParams {
	raw, err := ioutil.ReadFile("inaas.json")
	if err != nil {
		panic(err)
	}

	var params AppParams
	err = json.Unmarshal(raw, &params)
	if err != nil {
		panic(err)
	}

	return params
}
