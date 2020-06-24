package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	countMap := map[string]int{}
	words := strings.Fields(text)
	for _, word := range words {
		word = strings.ToLower(word)
		value, ok := countMap[word]
		if !ok {
			countMap[word] = 1
		} else {
			countMap[word] = value + 1
		}
	}
	for key := range countMap {
		fmt.Println(key, ":", countMap[key])
	}

}
