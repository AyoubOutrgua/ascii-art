package function

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArtTable() map[int]string {
	fileData, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error:", err)
		return map[int]string{}
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
	return asciiArtTable
}
