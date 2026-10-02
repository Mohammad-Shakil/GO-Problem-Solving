package main

import "fmt"

func findSame(a, b []int) []int {
	var result []int
	for _, val1 := range a {
		for _, val2 := range b {
			if val1 == val2 {
				result = append(result, val1)
			}
		}
	}
	return result
}

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}

	fmt.Println(findSame(a, b))
}
