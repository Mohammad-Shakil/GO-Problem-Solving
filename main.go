package main

import (
	"fmt"
	"strings"
)

func findlargest(line string) string {

	wordsofLine := strings.Fields(line)
	lar := 0
	var largest string
	for _, word := range wordsofLine {

		if len(word) > lar {
			lar = len(word)
			largest = word
		}

	}
	return largest
}
func main() {

	line := "I love programming in Go"

	fmt.Println("Largest:", findlargest(line))
}
