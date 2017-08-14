package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int) //init map

	str := strings.Fields(s) //breaks up a string based on spacing

	for _, word := range str {
		count[word]++
	}

	return count
}

func main() {
	wc.Test(WordCount)
}