package function

import (
	"fmt"
	"strings"
)
// a function to print the ascii values line by line
func PrintArt(artSlice [][]string) {
	for i := 0; i < 8; i++ {
		lineParts := []string{}
		for _, line := range artSlice {
			lineParts = append(lineParts, line[i])
		}
		fmt.Println(strings.Join(lineParts, " "))
	}
}
