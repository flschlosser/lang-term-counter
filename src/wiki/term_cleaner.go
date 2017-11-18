package wiki

import (
    "regexp"
    "log"
    "strings"
)

func SplitAndCountTerms(content string, termFrequency *TermFrequency) (totalTermCount int) {
    reg, err := regexp.Compile("[^a-zA-Z]+")
    if err != nil {
        log.Fatal(err)
    }
    cleanContent := reg.ReplaceAllString(content, " ")
    terms := strings.Split(cleanContent, " ")

    for _, term := range terms  {
        term = strings.ToLower(term)
        if len(term) > 0 {
            termFrequency.Inc(term)
        }
    }
    return len(terms)
}


