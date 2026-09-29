package main

import "fmt"

func filter(list []string, target int) []string {

	var new []string

	for _, value := range list {

		letter := []byte(value)
		if len(letter) > target {

			new = append(new, string(letter))
		}
	}
	return new
}

func main() {

	list := []string{"apple", "banana", "watermelon", "mango", "pineapple"}
	target := 5
	list = filter(list, target)
	fmt.Println(list)
}
