package main

import (
    "net/http"
    "fmt"
    "lang-term-counter/src/wiki"
)

type Collector struct {
    langCode string
    url string
    goalTermCount int
}

func (c *Collector) CollectTermFrequency() *wiki.TermFrequency  {
    var currentTermCount = 0

    termFrequency := wiki.NewTermFrequency()

    for currentTermCount < c.goalTermCount {
        response, err := http.Get(c.url)
        if err != nil {
            fmt.Printf("Error: unable to load content of url '%s'\n", c.url)
            return termFrequency
        }

        parser := wiki.NewContentParser()
        content := parser.Parse(response.Body)
        currentTermCount += wiki.SplitAndCountTerms(content, termFrequency)
    }
    return termFrequency
}
