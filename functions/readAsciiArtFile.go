package functions

import (
	"fmt"
	"os"
	"strings"
)

func ReadAsciiArtFile() map[int]string {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("Error:", err)
		return map[int]string{}
	}
	lines := strings.Split(string(file), "\n\n")
	// fmt.Println(lines)
	// fmt.Println(len(lines)/9)
	asciiArtTable := make(map[int]string)
	key := 32
	for _, line := range lines {
		asciiArtTable[key] = line
		key++
	}
	return asciiArtTable
}
