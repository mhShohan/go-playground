package main

import "fmt"

func main() {
	type Person struct {
		name      string
		age       int
		isStudent bool
	}

	var p1 = Person{name: "shohan", age: 23, isStudent: true}
	fmt.Println(p1)
}
