package function

import (
	"fmt"
	"strings"
)

func AppendArt(wordsSlice []string, asciiArtTable map[int]string) {
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
