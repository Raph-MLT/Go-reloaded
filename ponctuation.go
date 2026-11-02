package main

import "strings"

const punctuation = ".,!?:;"

func fixPunctuation(words []string) []string {
	result := []string{}

	for _, word := range words {
		rest := strings.TrimLeft(word, punctuation)
		start := strings.TrimSuffix(word, rest)

		if start != "" {
			if len(result) > 0 {
				result[len(result)-1] += start
			} else {
				result = append(result, start)
			}
		}

		if rest != "" {
			result = append(result, rest)
		}
	}

	return result
}
