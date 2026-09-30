package main

import "fmt"

func main() {

	makeSlice := make([]int, 0, 10)

	fmt.Println("Length:", len(makeSlice))
	fmt.Println("Capacity:", cap(makeSlice))

	fmt.Println("----------")
	makeSlice = append(makeSlice, 3)
	fmt.Println(makeSlice)

	fmt.Println("Length:", len(makeSlice))
	fmt.Println("Capacity:", cap(makeSlice))
}
