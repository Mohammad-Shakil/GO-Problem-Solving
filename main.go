package main

import "fmt"

func checkDuplicate(list []int) bool {

	seen := make(map[int]bool)

	for i := 0; i < len(list); i++ {
		if seen[list[i]] {
			return true
		} else {
			seen[list[i]] = true
		}
	}
	return false
}

func main() {

	list := []int{1, 2, 1, 3, 4, 5}

	fmt.Println(checkDuplicate(list))

}
