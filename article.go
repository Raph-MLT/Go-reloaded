package main

import "strings"

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			if startsWithVowelOrH(words[i+1]) {
				words[i] = words[i] + "n"
			}
		}
	}
	return words
}

func startsWithVowelOrH(word string) bool {
	word = strings.TrimLeft(word, "'")
	if word == "" {
		return false
	}
	first := strings.ToLower(word[:1])
	return strings.Contains("aeiouh", first)
}
