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
	function.AppendArt(strings.Split(os.Args[1], "\\n"), function.AsciiArtTable())
}
