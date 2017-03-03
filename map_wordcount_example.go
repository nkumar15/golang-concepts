package main

import (
	"fmt"
	"strings"
)

func WordCount(sentence string) map[string]int {
	words := strings.Fields(sentence)

	wordCountMap := make(map[string]int)

	for _, w := range words {
		count := wordCountMap[w]
		wordCountMap[w] = count + 1
	}
	return wordCountMap
}

func main() {
	wordCount := WordCount("That that is is that that is not is not. Is that it? It is.")
	for word, count := range wordCount {
		fmt.Printf("%s       : %d\n",word,count)
	}

}
