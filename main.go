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
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	blocks := strings.Split(string(data), "\n\n")
	// if blocks[0] {

	// }
	if strings.Count(blocks[0], "\n") == 8 {
		blocks[0] = blocks[0][1:]
	}
	art_table := make(map[int]string)
	key := 32
	for _, block := range blocks {
		art_table[key] = block
		key++
	}
	slice := [][]string{}
	sliceValue := []string{}
	for i := 0; i < len(inputText); i++ {
		for key, value := range art_table {
			if i+1 < len(inputText) && inputText[i] == '\\' && inputText[i+1] == 'n' {
				continue
			}
			if inputText[i] == byte(key) {
				sliceValue = strings.Split(value, "\n")
				slice = append(slice, sliceValue)
				sliceValue = []string{}
			}
		}
		if i+1 < len(inputText) && inputText[i] == '\\' && inputText[i+1] == 'n' {
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")
			sliceValue = append(sliceValue, "\n")

			slice = append(slice, sliceValue)
			i++
		}
	}

	// fmt.Println(slice)

	for i := 0; i < 8; i++ {
		lineParts := []string{}
		for _, line := range slice {
			lineParts = append(lineParts, line[i])
		}
		fmt.Println(strings.Join(lineParts, " "))
	}
}
