package main

import (
	"fmt"
	"os"
)


func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . <fichier d'entrée> <fichier de sortie>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFile, content, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		os.Exit(1)
	}
}