package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var ptr *int // pointer variable declaration
	fmt.Println("Value of ptr: ", ptr)

	num := 10       // num is an integer variable
	var ptr2 = &num // &num is the address of num

	fmt.Println("address of actual pointer: ", ptr2)
	fmt.Println("Value of actual pointer: ", *ptr2)
}
