package main

import (
	"fmt"
	"os"

	"asciiart/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid args!")
		return
	}
	inputText := os.Args[1]

	asciiArtTable := functions.ReadAsciiArtFile()

	for _, litter := range inputText {
		for key, art := range asciiArtTable {
			// if i+1 < len(inputText) && litter == '\\' && inputText[i+1] == 'n' {
			// 	fmt.Println()
			// 	continue
			// }
			if int(litter) == key {
				fmt.Print(art)
			}
		}
	}
}
