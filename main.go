package main

import "fmt"

func main() {

	list := []int{1, 2, 4, 5, 6}

	p1 := 6
	p2 := 6 + 1
	p3 := p1 * p2
	value := p3 / 2

	var sum int
	for _, val := range list {
		sum = sum + val
	}

	fmt.Println(value - sum)
}
