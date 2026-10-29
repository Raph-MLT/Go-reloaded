package main

import "strings"

func process(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		lines[i] = processLine(line)
	}

	return strings.Join(lines, "\n")
}

func processLine(line string) string {
	words := splitWords(line)
	return strings.Join(words, " ")
}

func splitWords(line string) []string {
	line = strings.ReplaceAll(line, "(up, ", "(up,")
	line = strings.ReplaceAll(line, "(low, ", "(low,")
	line = strings.ReplaceAll(line, "(cap, ", "(cap,")

	return strings.Fields(line)
}
