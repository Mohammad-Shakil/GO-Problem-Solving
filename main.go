package main

import "fmt"

func findSum(list []int, target int) (int, int) {

	var num1 int
	var num2 int

	for i := 0; i < len(list); i++ {

		for j := 1; j < len(list); j++ {

			if list[i]+list[j] == target {
				num1 = i
				num2 = j
			}
		}
	}
	return num1, num2
}

func main() {

	list := []int{2, 7, 11, 15}

	target := 9

	val1, val2 := findSum(list, target)

	fmt.Println(list[val1], list[val2])
}
