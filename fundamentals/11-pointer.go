package main

import "fmt"

func main() {
	var a = 10
	var b *int = &a

	*b = 15
	// a = 300

	fmt.Println(a, *b)
}
