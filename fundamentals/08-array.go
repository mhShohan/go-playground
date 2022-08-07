package main

import "fmt"

func main() {
	type Person struct {
		name      string
		age       int
		isStudent bool
	}

	//fixed array length
	students := [5]Person{}

	students[0] = Person{name: "shohan", age: 23, isStudent: true}

	fmt.Println(students)
}
