package main

import "fmt"

func countVowels(word string) int {

	count := 0

	for _, runes := range word {
		value := string(runes)

		switch value {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			count++
		}
	}
	return count
}

func main() {

	word := "Hello Worldee"

	fmt.Println("Total Vowel count:", countVowels(word))

}
