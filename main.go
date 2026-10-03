package main

import "fmt"

func findSecondSmallest(list []int) (int, int) {

	smallest := list[0]
	sesmallest := list[0]
	for i := 0; i < len(list); i++ {

		if list[i] < smallest {
			sesmallest = smallest
			smallest = list[i]
		}
		if list[i] > smallest && list[i] < sesmallest {
			sesmallest = list[i]
		}

	}

	return smallest, sesmallest
}

func main() {

	list := []int{10, 5, 20, 8, 15}

	smallest, secsmallest := findSecondSmallest(list)

	fmt.Printf("smallest number is: %d \nSecond smallest number is: %d", smallest, secsmallest)
}
