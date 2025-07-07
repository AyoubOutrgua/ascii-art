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
	inputSlice := []string{}
	str := ""
	for i := 0; i < len(inputText); i++ {
		// fmt.Println("str = ",str)
		if i+1 < len(inputText) && inputText[i] == '\\' && inputText[i+1] == 'n' {
			inputSlice = append(inputSlice, str)
			str = ""
			i++
		} else {
			str += string(inputText[i])
		}
	}
	// fmt.Println("str outside = ",str)
	inputSlice = append(inputSlice, str)
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	blocks := strings.Split(string(data), "\n\n")
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
	// fmt.Println(inputSlice)
	for _, v := range inputSlice {
		for i := 0; i < len(v); i++ {
			for key, value := range art_table {
				if v[i] == byte(key) {
					sliceValue = strings.Split(value, "\n")
					slice = append(slice, sliceValue)
					sliceValue = []string{}
				}
			}
		}
		PrintSlice(slice)
		slice = [][]string{}
		fmt.Println()
	}
}

func PrintSlice(slice [][]string) {
	for i := 0; i < 8; i++ {

		lineParts := []string{}
		for _, line := range slice {
			lineParts = append(lineParts, line[i])
		}
		lnSlice := CleanSlice(lineParts)
		// fmt.Println("len linepqrts : ", len(lnSlice))
		if len(lnSlice) != 0 {
			fmt.Println(strings.Join(lineParts, " "))
		}
		
	}
}

func CleanSlice(slice []string) []string {
	str := strings.Join(slice, " ")
	slice = strings.Fields(str)
	return slice
}
