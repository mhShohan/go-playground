package main

import "fmt"

func multipleReturnValue() (string, string) {
	return "mh", "shohan"
}

func main() {
	fName, lName := multipleReturnValue()

	fmt.Println(fName, lName)
}
