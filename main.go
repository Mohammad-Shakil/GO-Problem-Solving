package main

import "fmt"

func main() {

	list := []int{10, 5, 20, 8, 20, 15}

	max := list[0]
	var second int
	for i := 1; i < len(list); i++ {

		if list[i] > max {
			second = max
			max = list[i]
		}
		if list[i] > second && list[i] < max {
			second = list[i]
		}

	}

	fmt.Println("Max:", max, "Second max", second)
}
