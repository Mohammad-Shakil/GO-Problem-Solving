package main

import "fmt"

func chekDuplicate(list []int) []int {
	var newlist []int
	for i := 0; i < len(list); i++ {
		seen := false

		for j := 0; j < len(newlist); j++ {
			if list[i] == newlist[j] {
				seen = true
				break
			}
		}
		if seen == false {
			newlist = append(newlist, list[i])
		}
	}
	return newlist
}

func main() {

	list := []int{1, 2, 3, 2, 5}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Problem", chekDuplicate(list))
	fmt.Println("Correct", chekDuplicate(numbers))
}
