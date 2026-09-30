package main

import (
	"fmt"
	"strings"
)

func countWords(list string) int {

	slice := strings.Fields(list)

	return len(slice)
}

func main() {

	list := "I love Go programming"

	fmt.Println(countWords(list))
}
