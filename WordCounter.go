package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	stringArr := strings.Split(s, " ")
	for _, value := range stringArr {
		counter[value] += 1
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
