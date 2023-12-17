package main

import "fmt"

func main() {
	var name string
	name = "Budiono Siregar"
	fmt.Println(name)

	for i := 0; i < len(name); i++ {
		fmt.Println(name)
	}
}
