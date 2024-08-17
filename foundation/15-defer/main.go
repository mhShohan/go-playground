package main

import "fmt"

func main() {
	fmt.Println("--------Defer---------")

	// LIFO - Last in First Out

	defer fmt.Println("1")
	defer fmt.Println("2")

	fmt.Println("3")
	fmt.Println("4")

	defer fmt.Println("5")

	fmt.Println("6")
}
