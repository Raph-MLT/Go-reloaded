package main

import "strings"

func fixQuotes(words []string) []string {
	result := []string{}
	open := false

	for i := 0; i < len(words); i++ {
		word := words[i]

		if strings.TrimRight(word, punctuation) != "'" {
			result = append(result, word)
			continue
		}

		if !open {
			if i+1 < len(words) {
				words[i+1] = word + words[i+1]
			} else {
				result = append(result, word)
			}
			open = true
		} else {
			if len(result) > 0 {
				result[len(result)-1] += word
			} else {
				result = append(result, word)
			}
			open = false
		}
	}

	return result
}
