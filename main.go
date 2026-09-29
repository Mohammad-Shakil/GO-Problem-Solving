package main

import "fmt"

func freqChecker(list []int) map[int]int {

	newlist := make(map[int]int)

	for _, value := range list {
		newlist[value]++
	}
	return newlist
}

func main() {

	list := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 1, 2, 3, 6, 8, 8, 9, 9}

	newlist := freqChecker(list)
	for key, value := range newlist {
		fmt.Printf("\n%d : %d", key, value)
	}
}
