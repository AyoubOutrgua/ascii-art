package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid args!")
		return
	}
	// text := os.Args[1]
}
