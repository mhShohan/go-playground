package main

import (
	"fmt"
)

func main() {
	person := make(map[string]string)

	//assign value
	person["name"] = "mhShohan"
	person["school"] = "BSMRSTU"

	//reading value
	fmt.Println(person["name"])

	//delete value
	delete(person, "school")

	fmt.Println(person)
}
