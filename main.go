package main

import "fmt"

func revers(word string) bool {

	list := []rune(word)

	low := 0
	high := len(list) - 1

	for low < high {
		list[low], list[high] = list[high], list[low]
		low++
		high--
	}

	final := string(list)
	return final == word
}

func main() {

	list := []string{"shakil", "uzzal", "lil", "pop"}

	for _, value := range list {

		truee := revers(value)

		if truee {
			fmt.Println("\n", value, "is a palindrome")
		} else {
			fmt.Println("\n", value, "is not a palindrome")
		}
	}
}
