// +build ignore

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

var wordsCount map[string]int

func WordCount(s string) map[string]int {
	wordsCount = make(map[string]int)
	words := strings.Fields(s)

	for _, v := range words {
		_, ok := wordsCount[v]

		if !ok {
			wordsCount[v] = 1
		} else {
			wordsCount[v]++
		}
	}

	return wordsCount
}

func main() {
	wc.Test(WordCount)
}
