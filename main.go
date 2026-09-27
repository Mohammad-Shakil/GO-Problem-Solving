package main

import "fmt"

func checkfreq(list []int) map[int]int {

	freq := make(map[int]int)

	for _, value := range list {

		freq[value]++

	}

	return freq
}

func main() {

	list := []int{1, 1, 1, 2, 2, 3, 1, 4, 2, 5, 6, 7, 5, 3, 4, 6, 6}

	res := checkfreq(list)

	for num, count := range res {
		fmt.Printf("\n%d : %d bar\n", num, count)
	}
}
