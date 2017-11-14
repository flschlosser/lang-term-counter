package main

import (
	"io"
	"golang.org/x/net/html"
	"bytes"
)

type HtmlParser struct {}

func (s *HtmlParser) Parse(reader io.Reader) (content string, err error) {

	tokenizer := html.NewTokenizer(reader)
	printNextTextToken := false
	withinContent := false

	// to count the divs within the content div
	countDivs := false
	divCount := 0

	var buffer bytes.Buffer

	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.ErrorToken:
			return buffer.String(), nil // EOF

		case tokenType == html.StartTagToken:
			token := tokenizer.Token()

			if token.Data == "div" && countDivs {
				divCount++
			}

			if token.Data == "div" {
				for _, attribute := range token.Attr {
					if attribute.Key == "id" && attribute.Val == "mw-content-text" {
						withinContent = true
						countDivs = true
						break
					}
				}
			}

			if token.Data == "p" {
				printNextTextToken = withinContent
			}

		case tokenType == html.EndTagToken:
			token := tokenizer.Token()

			if token.Data == "div" && countDivs {
				divCount--
				if divCount == 0 {
					countDivs = false
				}
			}

			if token.Data == "p" {
				printNextTextToken = false
			}

		case tokenType == html.TextToken:
			if printNextTextToken {
				token := tokenizer.Token()
				buffer.WriteString(token.Data)
			}
		}
	}
	return buffer.String(), nil
}
