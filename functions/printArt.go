package function

import (
	"fmt"
	"strings"
)

func PrintArt(artSlice [][]string) {
	for i := 0; i < 8; i++ {
		lineParts := []string{}
		for _, line := range artSlice {
			lineParts = append(lineParts, line[i])
		}
		fmt.Println(strings.Join(lineParts, " "))
	}
}
