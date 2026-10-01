package main

import "fmt"

func chekDuplicate(list []int) bool {
	seen := false
	for i := 0; i < len(list); i++ {

		for j := i + 1; j < len(list); j++ {
			if list[i] == list[j] {
				seen = true
				return seen
			}

		}
	}

	return seen
}

func main() {

	list := []int{1, 2, 3, 2, 5}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(list, "Duplicate result:", chekDuplicate(list))
	fmt.Println(numbers, "Duplicate result:", chekDuplicate(numbers))
}
