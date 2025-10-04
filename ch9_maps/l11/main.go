package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	// ?
	freqMap := make(map[string]int)

	for _, msg := range messages {
		words := strings.Split(msg, " ")

		for _, word := range words {
			if word == "" {
				continue
			}
			ignoredCase := strings.ToLower(word)
			freqMap[ignoredCase] += 1
		}
	}

	return len(freqMap)
}
