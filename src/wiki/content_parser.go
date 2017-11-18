package wiki

import (
    "io"
    "golang.org/x/net/html"
    "bytes"
)

const (
    contentDivIdValue = "mw-content-text"
)

type contentParser struct {
    shouldPrintTextOfToken bool
    divCount               int
    buffer                 bytes.Buffer
}

func NewContentParser() *contentParser {
    parser := &contentParser{
        shouldPrintTextOfToken: false,
        divCount:               0,
        buffer:                 bytes.Buffer{},
    }
    return parser
}

func (p *contentParser) Parse(reader io.Reader) string {
    tokenizer := html.NewTokenizer(reader)

    p.shouldPrintTextOfToken = false
    p.divCount = 0
    p.buffer = bytes.Buffer{}

    for {
        tokenType := tokenizer.Next()
        token := tokenizer.Token()

        switch tokenType {
        case html.ErrorToken:
            return p.buffer.String()

        case html.StartTagToken:
            p.handleStartTagToken(token)

        case html.TextToken:
            p.handleTextToken(token)

        case html.EndTagToken:
            p.handleEndTagToken(token)

        }
    }
    return p.buffer.String()
}

func (p *contentParser) handleStartTagToken(token html.Token) {
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

func (p *contentParser) handleEndTagToken(token html.Token) {
    if token.Data == "div" && p.divCount > 0 {
        p.divCount--
    } else
    if token.Data == "p" {
        p.shouldPrintTextOfToken = false
    }
}

func (p *contentParser) handleTextToken(token html.Token) {
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
