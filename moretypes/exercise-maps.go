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
		if _, ok := wordsCount[v]; ok {
			wordsCount[v]++
		} else {
			wordsCount[v] = 1
		}
	}

	return wordsCount
}

func main() {
	wc.Test(WordCount)
}
