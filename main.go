package main

import (
	"fmt"
	"os"
	"strings"

	function "asciiart/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("number of arguments is wrong!")
		return
	}
	for _, v := range os.Args[1] {
		if v < 32 || v >= 127 {
			fmt.Println("error, string is not valid ")
			return
		}
	}
	function.AppendArt(strings.Split(os.Args[1], "\\n"), function.AsciiArtTable())
}
