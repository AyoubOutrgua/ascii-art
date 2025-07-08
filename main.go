package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("number of arguments is wrong!")
		return
	}
	inputText := os.Args[1]
	wordsSlice := strings.Split(inputText, "\\n")
	fileData, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	blocks := strings.Split(string(fileData), "\n\n")
	if strings.Count(blocks[0], "\n") == 8 {
		blocks[0] = blocks[0][1:]
	}
	asciiArtTable := make(map[int]string)
	key := 32
	for _, block := range blocks {
		asciiArtTable[key] = block
		key++
	}
	artSlice := [][]string{}
	for _, word := range wordsSlice {
		if len(word) == 0 {
			fmt.Println()
			continue
		}
		for _, letter := range word {
			for key, value := range asciiArtTable {
				if letter == rune(key) {
					artSlice = append(artSlice, strings.Split(value, "\n"))
				}
			}
		}
		PrintArt(artSlice)
		artSlice = [][]string{}
	}
}

func PrintArt(artSlice [][]string) {
	for i := 0; i < 8; i++ {
		lineParts := []string{}
		for _, line := range artSlice {
			lineParts = append(lineParts, line[i])
		}
		fmt.Println(strings.Join(lineParts, " "))
	}
}
