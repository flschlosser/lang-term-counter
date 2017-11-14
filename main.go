package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://en.wikipedia.org/wiki/Hans_Mork"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: unable to load content of url '%s'\n", url)
		return
	}

	parser := HtmlParser{}
	content, err := parser.Parse(response.Body)
	if err != nil {
		fmt.Printf("Error unable to parse content because error:\n%s\n", err)
	}
	fmt.Printf("Content:\n%s\n", content)
}