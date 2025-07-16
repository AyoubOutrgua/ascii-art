package main

import (
	"fmt"
	"os"

	function "asciiart/functions"
)

func main() {
	// checking the validity of input
	if len(os.Args) != 2 {
		fmt.Println("number of arguments is wrong!")
		return
	}

	// check range of printable characters
	inputText := os.Args[1]
	if !function.CheckRange(inputText) {
		fmt.Println("error, string is not valid ")
		return
	}

	// split input text to slice of string
	wordsSlice := function.SplitInput(inputText)

	// calling functions to handle the input
	function.AppendArt(wordsSlice, function.AsciiArtTable())
}
