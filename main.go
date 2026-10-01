package main

import "fmt"

func moveZero(list []int) []int {

	var nlist []int
	var zeros int
	for _, value := range list {

		if value != 0 {
			nlist = append(nlist, value)
		} else {
			zeros++
		}
	}
	for i := 1; i <= zeros; i++ {
		nlist = append(nlist, 0)
	}
	return nlist
}

func main() {

	list := []int{0, 1, 0, 3, 12}

	fmt.Println(moveZero(list))
}
