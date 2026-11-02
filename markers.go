package main

import (
	"strconv"
	"strings"
)

func applyMarkers(words []string) []string {
	result := []string{}

	for _, word := range words {
		action, count, isMarker := readMarker(word)

		if !isMarker {
			result = append(result, word)
			continue
		}

		start := len(result) - count
		if start < 0 {
			start = 0
		}
		for i := start; i < len(result); i++ {
			result[i] = transform(result[i], action)
		}
	}

	return result
}

func readMarker(word string) (string, int, bool) {

	switch word {
	case "(hex)":
		return "hex", 1, true
	case "(bin)":
		return "bin", 1, true
	case "(up)":
		return "up", 1, true
	case "(low)":
		return "low", 1, true
	case "(cap)":
		return "cap", 1, true
	}

	for _, action := range []string{"up", "low", "cap"} {
		prefix := "(" + action + ","

		if strings.HasPrefix(word, prefix) && strings.HasSuffix(word, ")") {
			numberText := strings.TrimPrefix(word, prefix)
			numberText = strings.TrimSuffix(numberText, ")")

			count, err := strconv.Atoi(numberText)
			if err == nil && count > 0 {
				return action, count, true
			}
		}
	}

	return "", 0, false
}

func transform(word string, action string) string {
	switch action {
	case "hex":
		return toDecimal(word, 16)
	case "bin":
		return toDecimal(word, 2)
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		return capitalize(word)
	}
	return word
}

func toDecimal(word string, base int) string {
	number, err := strconv.ParseInt(word, base, 64)
	if err != nil {

		return word
	}
	return strconv.FormatInt(number, 10)
}

func capitalize(word string) string {
	if word == "" {
		return word
	}
	lower := strings.ToLower(word)
	return strings.ToUpper(lower[:1]) + lower[1:]
}
