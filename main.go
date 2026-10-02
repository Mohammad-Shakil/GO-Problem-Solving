package main

import (
	"fmt"
	"strings"
)

func reversSentece(line string) string {

	lientowords := strings.Fields(line)
	var finallist []string
	for _, word := range lientowords {
		sliceWord := []rune(word)

		low := 0
		high := len(sliceWord) - 1
		for low < high {

			sliceWord[low], sliceWord[high] = sliceWord[high], sliceWord[low]
			low++
			high--
		}
		wordback := string(sliceWord)
		finallist = append(finallist, wordback)
	}
	final := strings.Join(finallist, " ")
	return final

}

func main() {

	line := "Hello how are you"

	fmt.Println(reversSentece(line))
}
