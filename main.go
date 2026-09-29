package main

import "fmt"

func stringReversal(p []string) []string {
	var newlist []string
	for _, element := range p {
		letters := []rune(element)

		low := 0
		high := len(letters) - 1
		for low < high {
			letters[low], letters[high] = letters[high], letters[low]
			low++
			high--
		}
		bToS := string(letters)
		newlist = append(newlist, bToS)
	}
	return newlist
}

func main() {

	list := []string{"shakil", "kamal", "Parul", "Fahad", "uzzal", "sumaiya", "prantu", "pudding", "chiku"}

	revList := stringReversal(list)

	for _, element := range revList {
		fmt.Printf("\n Name: %s", element)
	}

}
