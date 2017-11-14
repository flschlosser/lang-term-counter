package main

import (
	"io"
	"golang.org/x/net/html"
	"bytes"
)

type HtmlParser struct {
	printNextTextToken bool
	withinContent bool

	// to count the divs within the content div
	countDivs bool
	divCount int

	buffer bytes.Buffer
}

func (p *HtmlParser) Parse(reader io.Reader) (content string, err error) {

	tokenizer := html.NewTokenizer(reader)
	p.printNextTextToken = false
	p.withinContent = false
	p.countDivs = false
	p.divCount = 0

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		switch {
		case tokenType == html.ErrorToken: // EOF
			return p.buffer.String(), nil

		case tokenType == html.StartTagToken:
			p.handleStartTagToken(token)

		case tokenType == html.EndTagToken:
			p.handleEndTagToken(token)

		case tokenType == html.TextToken:
			p.handleTextToken(token)
		}
	}
	return p.buffer.String(), nil
}

func (p *HtmlParser) handleStartTagToken(token html.Token) {
	if token.Data == "div" && p.countDivs {
		p.divCount++
	}
	if token.Data == "div" {
		for _, attribute := range token.Attr {
			if attribute.Key == "id" && attribute.Val == "mw-content-text" {
				p.withinContent = true
				p.countDivs = true
				break
			}
		}
	}
	if token.Data == "p" {
		p.printNextTextToken = p.withinContent
	}
}

func (p *HtmlParser) handleEndTagToken(token html.Token) {
	if token.Data == "div" && p.countDivs {
		p.divCount--
		if p.divCount == 0 {
			p.countDivs = false
		}
	}
	if token.Data == "p" {
		p.printNextTextToken = false
	}
}

func (p *HtmlParser) handleTextToken(token html.Token) {
	if p.printNextTextToken {
		p.buffer.WriteString(token.Data)
	}
}
