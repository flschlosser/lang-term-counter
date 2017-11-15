package main

import (
	"io"
	"golang.org/x/net/html"
	"bytes"
)

const (
	contentDivIdValue = "mw-content-text"
)

type WikiContentParser struct {
	shouldPrintTextOfToken bool
	divCount               int
	buffer                 bytes.Buffer
}

func (p *WikiContentParser) Parse(reader io.Reader) string {
	tokenizer := html.NewTokenizer(reader)

	p.shouldPrintTextOfToken = false
	p.divCount = 0
	p.buffer = bytes.Buffer{}

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		switch {
		case tokenType == html.ErrorToken: // EOF
			return p.buffer.String()

		case tokenType == html.StartTagToken:
			p.handleStartTagToken(token)

		case tokenType == html.EndTagToken:
			p.handleEndTagToken(token)

		case tokenType == html.TextToken:
			p.handleTextToken(token)
		}
	}
	return p.buffer.String()
}

func (p *WikiContentParser) handleStartTagToken(token html.Token) {
	// to handle divs within the content div
	if token.Data == "div" && p.divCount > 0 {
		p.divCount++
	} else
	if token.Data == "p" {
		p.shouldPrintTextOfToken = p.divCount > 0
	} else
	if isTokenTheContentDiv(token) {
		p.divCount++
	}
}

func (p *WikiContentParser) handleEndTagToken(token html.Token) {
	if token.Data == "div" && p.divCount > 0 {
		p.divCount--
	}
	if token.Data == "p" {
		p.shouldPrintTextOfToken = false
	}
}

func (p *WikiContentParser) handleTextToken(token html.Token) {
	if p.shouldPrintTextOfToken {
		p.buffer.WriteString(token.Data)
	}
}

func isTokenTheContentDiv(token html.Token) bool {
	if token.Data != "div" {
		return false
	}
	for _, attribute := range token.Attr {
		if attribute.Key == "id" && attribute.Val == contentDivIdValue {
			return true
		}
	}
	return false
}
