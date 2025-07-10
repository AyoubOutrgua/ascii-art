package main

import (
	"fmt"
	"os"
	"strings"

	function "asciiart/functions"
)

func main() {
	// checking the validity of input
	if len(os.Args) != 2 {
		fmt.Println("number of arguments is wrong!")
		return
	}
	for _, v := range os.Args[1] {
		if v < 32 || v > 126 {
			fmt.Println("error, string is not valid ")
			return
		}
	}
	// calling functions to handle the input
	wordsSlice := strings.Split(os.Args[1], "\\n")
	if len(wordsSlice[0]) == 0 {
		wordsSlice = wordsSlice[1:]
	}
	function.AppendArt(wordsSlice, function.AsciiArtTable())
}
