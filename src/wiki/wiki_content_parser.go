package wiki

import (
	"io"
	"golang.org/x/net/html"
	"bytes"
)

const (
	contentDivIdValue = "mw-content-text"
)

type wikiContentParser struct {
	shouldPrintTextOfToken bool
	divCount               int
	buffer                 bytes.Buffer
	handlers 			   map[html.TokenType]func(token html.Token)(finish bool)
}

func NewWikiContentParser() *wikiContentParser {
	parser := &wikiContentParser{
		shouldPrintTextOfToken: false,
		divCount:               0,
		buffer:                 bytes.Buffer{},
		handlers:				make(map[html.TokenType]func(token html.Token)(finish bool)),
	}
	parser.handlers[html.StartTagToken] = parser.handleStartTagToken
	parser.handlers[html.EndTagToken] = parser.handleEndTagToken
	parser.handlers[html.TextToken] = parser.handleTextToken
	parser.handlers[html.ErrorToken] = parser.handleErrorToken
	return parser
}

func (p *wikiContentParser) Parse(reader io.Reader) string {
	tokenizer := html.NewTokenizer(reader)

	p.shouldPrintTextOfToken = false
	p.divCount = 0
	p.buffer = bytes.Buffer{}

	for {
		if p.processToken(tokenizer.Next(), tokenizer.Token()) {
			break
		}
	}
	return p.buffer.String()
}

func (p *wikiContentParser) processToken(tokenType html.TokenType, token html.Token) (finish bool) {
	handler := p.handlers[tokenType]
	if handler == nil {
		return false
	}
	return handler(token)
}

func (p *wikiContentParser) handleStartTagToken(token html.Token) (finish bool) {
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
	return false
}

func (p *wikiContentParser) handleEndTagToken(token html.Token) (finish bool) {
	if token.Data == "div" && p.divCount > 0 {
		p.divCount--
	} else
	if token.Data == "p" {
		p.shouldPrintTextOfToken = false
	}
	return false
}

func (p *wikiContentParser) handleTextToken(token html.Token) (finish bool) {
	if p.shouldPrintTextOfToken {
		p.buffer.WriteString(token.Data)
	}
	return false
}

func (p *wikiContentParser) handleErrorToken(token html.Token) (finish bool) {
	return true // EOF
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
