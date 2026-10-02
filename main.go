package main

import "fmt"

func findSame(a, b []int) ([]int, map[int]bool) {

	register := make(map[int]bool)
	var res []int
	for _, val := range a {
		register[val] = true
	}

	for _, value := range b {
		if register[value] {
			res = append(res, value)
		}
	}
	return res, register
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}

	fmt.Println(findSame(a, b))

}
