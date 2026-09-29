package main

import "fmt"

func removeDuplicate(list []int) []int {

	var newlist []int

	for _, value := range list {
		seen := false
		for _, val := range newlist {

			if value == val {
				seen = true
				break
			}
		}
		if seen != true {
			newlist = append(newlist, value)
		}
	}
	return newlist
}

func main() {

	numbers := []int{1, 2, 2, 3, 4, 4, 5}

	fmt.Println(removeDuplicate(numbers))

}
