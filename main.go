package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func name(p string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name:")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	fmt.Println("Hello", name)
}
func main() {

	var name string

	fmt.Print("Enter name:")
	fmt.Scanln(&name)
	fmt.Println(name)
}
