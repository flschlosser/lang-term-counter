package main

import (
    "fmt"
)

func main() {
    enCollector := Collector{
        "en",
        "https://en.wikipedia.org/wiki/Special:Random",
        1000}

    enTermFrequency := enCollector.CollectTermFrequency()

    fmt.Printf("english term frequency:\n%s\n", enTermFrequency)
}