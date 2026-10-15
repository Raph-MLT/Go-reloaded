package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	inputFile := os.Args[1]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	words := strings.Fields(string(content))

	result := strings.Join(words, " ")

	os.WriteFile(os.Args[2], []byte(result), 0644)
}