package function

import (
	"fmt"
	"strings"
)
// a function to range in the input and print it if matches with the map values
func AppendArt(wordsSlice []string, asciiArtTable map[int]string) {
	artSlice := [][]string{}
	//check if there is an empty string just print a newline and skip
	for _, word := range wordsSlice {
		if len(word) == 0 {
			fmt.Println()
			continue
		}
		/*range in each part letter by letter and check if it matches the map values
		then append it in a slice to print it*/
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
