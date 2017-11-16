package main

import (
	"fmt"
	"net/http"
	"lang-term-counter/src/wiki"
)

func main() {

	url := "https://en.wikipedia.org/wiki/Hans_Mork"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: unable to load content of url '%s'\n", url)
		return
	}

	parser := wiki.NewWikiContentParser()
	content := parser.Parse(response.Body)

	fmt.Printf("Content:\n%s\n", content)
}