package main

import (
	"fmt"
)

func main() {
	//Arithmetic  Operators ===>> +,-,*,/,%
	//Relation Operators ===>> ==, <,>,<=,>=,!=
	//Logical Operators  ====> &&,||, !

	name := "mhShohan"
	age := 23

	isValid := !(name == "shohan") && age > 18
	fmt.Println(isValid)
}