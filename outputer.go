package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Outputer interface {
	Output(metadata UrlMetadata)
}

type JsonOutputer struct{}
type GreppableOutputer struct{}

func (out GreppableOutputer) Output(metadata UrlMetadata) {
	fmt.Printf(metadata.Url)
	if metadata.StatusCode != 0 {
		fmt.Printf("\tStatusCode:%d", metadata.StatusCode)
	}
	for name, value := range metadata.ResponseHeasers {
		fmt.Printf("\t%s: %s", name, value)
	}
	fmt.Println()

	return
}

func (out JsonOutputer) Output(metadata UrlMetadata) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(metadata)

	return
}
