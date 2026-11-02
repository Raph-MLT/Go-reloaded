package main

import "strconv"

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
	}
	return "", 0, false
}

func transform(word string, action string) string {
	switch action {
	case "hex":
		return toDecimal(word, 16)
	case "bin":
		return toDecimal(word, 2)
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
