package main

import (
	"fmt"
)

func main() {

	var word string
	fmt.Print("Enter Text: ")
	fmt.Scanln(&word)

	listOfWord := []rune(word)

	low := 0
	high := len(listOfWord) - 1

	for i := 0; low < high; i++ {

		listOfWord[low], listOfWord[high] = listOfWord[high], listOfWord[low]
		low++
		high--

	}

	fmt.Println("\n", string(listOfWord))
}
