package main

import "fmt"

func main() {

	//string typed
	const firstName string = "Mehdi Hasan "
	var lastName string = "Shohan"

	fmt.Println(firstName + lastName)

	// number type
	const age int = 23
	const income float32 = 1000.34
	fmt.Println(age)

	//boolean typed
	const isAdult bool = true
	fmt.Println(isAdult)

	//compound variable
	const fName, lName, isStudent = "mh", "Shohan", true

	// instant variable declaration , only work in the function
	universityName := "BSMRSTU"
	fmt.Println(universityName)
}
