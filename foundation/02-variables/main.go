package main

import "fmt"

func main() {
	// Declare a variable and initialize it
	var username string = "MH Shohan"
	fmt.Println(username)
	fmt.Printf("Type of username: %T \n", username)

	var age int = 25
	fmt.Println(age)
	fmt.Printf("Type of age: %T \n", age)

	var isMarried bool = false
	fmt.Println(isMarried)
	fmt.Printf("Type of isMarried: %T \n", isMarried)

	// default value and some aliases
	var defaultString string
	var defaultInt int
	fmt.Println(defaultString)
	fmt.Println(defaultInt)

	// Short hand declaration
	// var name string = "MH Shohan"
	name := "MH Shohan"
	gpa := 3.925467456734545
	fmt.Println(name)
	fmt.Println(gpa)
}
